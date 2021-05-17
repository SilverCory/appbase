package loader

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/silvercory/appbase/config"
	"github.com/silvercory/appbase/migrations"
	"github.com/silvercory/appbase/server"

	"github.com/go-redis/redis/v8"

	"github.com/rs/zerolog"
)

// Loader does stuff
type Loader interface {
	// Setup an initial setup.
	Setup() error
	// Run should start it's own goroutine if it's blockingn
	Run() error
	// Get the instance.
	Get() interface{}
	// Close well... Closes...
	Close() error
}

var (
	logger  = zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()
	loaders = make(map[string]Loader)

	err error

	cfg       config.Config
	sqlDB     *sql.DB
	redisDB   *redis.Client
	webServer *server.Server
)

func AddLoader(n string, l Loader) {
	loaders[n] = l
}

func GetLoader(n string) interface{} {
	return loaders[n].Get()
}

func Setup() {
	if err := cfg.Load(); err != nil {
		logger.Error().Err(err).Msg("Loading configuration from env.")
		os.Exit(1)
		return
	}

	if cfg.MySQL.Enabled {
		if err = migrations.EnsureUpToDate(logger, cfg.MySQL); err != nil {
			logger.Error().Err(err).Msg("Migrating database failed!")
			os.Exit(1)
			return
		}

		sqlDB, err = sql.Open(cfg.MySQL.DatabaseType, cfg.MySQL.URI)
		if err != nil {
			logger.Error().Err(err).Msg("Opening sql connection.")
			os.Exit(1)
			return
		}
		if err := sqlDB.Ping(); err != nil {
			logger.Error().Err(err).Msg("Pinging sql DB.")
			os.Exit(1)
			return
		}
	}

	if cfg.Redis.Enabled {
		redisDB = redis.NewClient(&redis.Options{
			Addr:     cfg.Redis.Address,
			Password: cfg.Redis.Password,
			DB:       cfg.Redis.Database,
		})
		if err := redisDB.Ping(context.TODO()).Err(); err != nil {
			logger.Error().Err(err).Msg("Pinging redis DB.")
			os.Exit(1)
			return
		}
	}

	webServer, err = server.NewServer(logger, cfg.Web)
	if err != nil {
		logger.Error().Err(err).Msg("Setting up webserver")
		os.Exit(1)
		return
	}

	for k, v := range loaders {
		if err := v.Setup(); err != nil {
			logger.Error().Err(err).Msgf("Setting up loader %q", k)
			os.Exit(1)
			return
		}
	}
}

func Run() {
	go func() {
		if err := webServer.Start(cfg.Web.ListenAddress, false); err != nil {
			logger.Error().Err(err).Msg("Start web failed.")
		}
	}()

	for k, v := range loaders {
		if err := v.Run(); err != nil {
			logger.Error().Err(err).Msgf("Starting up loader %q", k)
			os.Exit(1)
			return
		}
	}

	_, _ = fmt.Fprintln(os.Stderr, "Waiting for interrupt... (Ctrl+C to stop)")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	if err := webServer.Close(); err != nil {
		logger.Error().Err(err).Msg("Closing webserver failed.")
		os.Exit(1)
	}

	for k, v := range loaders {
		if err := v.Close(); err != nil {
			logger.Error().Err(err).Msgf("Closing loader %q", k)
			os.Exit(1)
			return
		}
	}
}

func Logger() zerolog.Logger {
	return logger
}

func Config() *config.Config {
	return &cfg
}

func MySQL() *sql.DB {
	return sqlDB
}

func Redis() *redis.Client {
	return redisDB
}

func WebServer() *server.Server {
	return webServer
}

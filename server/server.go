package server

import (
	"fmt"
	"html/template"

	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
	"github.com/whipphotos/appbase/config"
	"github.com/whipphotos/appbase/easter_egg"
)

type Server struct {
	engine *gin.Engine

	logger   zerolog.Logger
	conf     config.Web
	launcher Launcher
	redisDB  *redis.Client

	closers []func() error
}

// Handler an interface for sections of the site that handle.
type Handler interface {
	RegisterHandlers() error
}

func NewServer(l zerolog.Logger, conf config.Web) (*Server, error) {
	var ret = &Server{
		engine: gin.New(),

		logger:   l,
		conf:     conf,
		launcher: NewLauncher(l, conf),
	}

	// Launcher closer.
	ret.closers = append(ret.closers, ret.launcher.Close)

	// Load the HTML templates
	// Templating
	ret.engine.SetFuncMap(template.FuncMap{
		"comments": func(s string) template.HTML { return template.HTML(s) },
		"ASCII":    easter_egg.GetAscii,
	})
	//ret.engine.LoadHTMLGlob(ret.conf.TemplateGlob)

	// Static files to load
	ret.engine.Use(
		gin.Recovery(),
		ginzerolog.Logger("gin"),
		//static.Serve("/", static.LocalFile(ret.conf.StaticFilePath, false)),
	)

	return ret, nil
}

func (s *Server) Start(listenAddr string, tls bool) error {
	var err error
	if !tls {
		err = s.engine.Run(listenAddr)
	} else {
		err = s.launcher.RunAutoTLS(s.engine)
	}

	if err != nil {
		return fmt.Errorf("web: server start: %w", err)
	}
	return nil
}

func (s *Server) Close() error {
	for _, v := range s.closers {
		if err := v(); err != nil {
			s.logger.Error().Err(err).Msg("Unable to close closer")
		}
	}

	_ = s.launcher.Close() // launcher Close doesn't actually return an error.
	return nil
}

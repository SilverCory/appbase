package config

import (
	"fmt"
	"net/url"
	"time"

	"github.com/caarlos0/env/v6"
)

// Config the main configuration.
type Config struct {
	MySQL   MySQL
	Redis   Redis
	Web     Web
	OwO     OwO
	Customs map[string]Parser
}

// Parser is an interface for custom configurations.
type Parser interface {
	// ParseEnv should call and return `return caarlos0.env.Parse(&this)`.
	ParseEnv() error
}

// Redis configures redis.
type Redis struct {
	Network  string `env:"REDIS_NETWORK" envDefault:"tcp"`
	Address  string `env:"REDIS_ADDRESS" envDefault:"127.0.0.1:6379"`
	Password string `env:"REDIS_PASSWORD" envDefault:"Hunter1"`
	Database int    `env:"REDIS_DATABASE" envDefault:"1"`
	Enabled  bool   `env:"REDIS_ENABLED" envDefault:"false"`
}

// MySQL configures MySQL.
type MySQL struct {
	DatabaseType string `env:"MYSQL_DATABASE_TYPE" envDefault:"mysql"`
	URI          string `env:"MYSQL_URI" envDefault:"username:password@tcp(127.0.0.1:3306)/database?charset=utf8&parseTime=True&loc=Local"`
	Enabled      bool   `env:"MYSQL_ENABLED" envDefault:"false"`
}

// Web configures gin and other web elements
type Web struct {
	//StaticFilePath   string   `env:"WEB_STATIC_FILE_PATH,file" envDefault:"./static/"`
	ListenAddress    string   `env:"WEB_LISTEN_ADDRESS" envDefault:":8080"`
	LogAuthKey       string   `env:"WEB_LOG_AUTH_KEY"`
	TemplateGlob     string   `env:"WEB_TEMPLATE_GLOB" envDefault:"templates/**/*.tmpl"`
	SentryDSN        string   `env:"WEB_SENTRY_DSN"`
	CSRFSecret       string   `env:"WEB_CSRF_SECRET"`
	CSPReportWebHook string   `env:"WEB_CSP_REPORT_WEBHOOK"`
	DomainNames      []string `env:"WEB_DOMAIN_NAMES" envSeparator:":" envDefault:"example.com:example.net"`
}

type OwO struct {
	UploadURL url.URL       `env:"OWO_UPLOAD_URL" envDefault:"https://api.awau.moe/upload/?key=uuid"`
	Timeout   time.Duration `env:"OWO_TIMEOUT" envDefault:"10s"`
	URL       url.URL       `env:"OWO_URL" envDefault:"https://awau.moe/"`
}

// Load loads the config.
func (c *Config) Load() error {
	if err := env.Parse(c); err != nil {
		return err
	}

	for k, v := range c.Customs {
		if err := v.ParseEnv(); err != nil {
			return fmt.Errorf("error calling ParseEnv for %q: %w", k, err)
		}
	}

	return nil
}

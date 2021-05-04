package server

import (
	"crypto/tls"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/silvercory/appbase/config"

	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog"

	"golang.org/x/crypto/acme/autocert"
)

type Launcher struct {
	logger zerolog.Logger
	conf   config.Web

	closers map[string]func() error
}

func NewLauncher(logger zerolog.Logger, webConf config.Web) Launcher {
	return Launcher{
		logger:  logger,
		conf:    webConf,
		closers: make(map[string]func() error),
	}
}

func (l *Launcher) RunAutoTLS(e *gin.Engine) error {
	m := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(l.conf.DomainNames...),
	}

	var dir = l.cacheDir()
	l.logger.Debug().Msgf("Using cache dir: %s", dir)
	if err := os.MkdirAll(dir, 0700); err != nil {
		l.logger.Warn().Msgf("warning: autocert.NewListener not using a cache: %v", err)
	} else {
		m.Cache = autocert.DirCache(dir)
	}

	go func() {
		var httpServer = l.getServer()
		httpServer.Addr = l.conf.ListenAddress + ":80"
		httpServer.Handler = m.HTTPHandler(nil)

		l.closers["http autocert server"] = httpServer.Close
		if err := httpServer.ListenAndServe(); err != nil {
			l.logger.Warn().Err(err).Msg("AutoCert server error.")
		}
	}()

	return l.runWithManager(e, m, l.conf.ListenAddress+":443")
}

func (l *Launcher) runWithManager(r http.Handler, m *autocert.Manager, address string) error {
	var httpServer = l.getServer()
	httpServer.Addr = address
	httpServer.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}
	httpServer.Handler = r

	l.closers["https server"] = httpServer.Close
	return httpServer.ListenAndServeTLS("", "")
}

func (l *Launcher) cacheDir() string {
	const base = "golang-autocert"
	switch runtime.GOOS {
	case "darwin":
		return filepath.Join(l.homeDir(), "Library", "Caches", base)
	case "windows":
		for _, ev := range []string{"APPDATA", "CSIDL_APPDATA", "TEMP", "TMP"} {
			if v := os.Getenv(ev); v != "" {
				return filepath.Join(v, base)
			}
		}
		// Worst case:
		return filepath.Join(l.homeDir(), base)
	}
	if xdg := os.Getenv("XDG_CACHE_HOME"); xdg != "" {
		return filepath.Join(xdg, base)
	}
	return filepath.Join(l.homeDir(), ".cache", base)
}

func (l *Launcher) homeDir() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
	}
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return "/"
}

func (l *Launcher) Close() error {
	for desc, closeFn := range l.closers {
		if err := closeFn(); err != nil {
			l.logger.Error().Err(err).Msgf("Unable to close %s", desc)
		}
	}
	return nil
}

func (l *Launcher) getServer() *http.Server {
	return &http.Server{
		ReadHeaderTimeout: 3 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      15 * time.Second,
		MaxHeaderBytes:    2048,
	}
}

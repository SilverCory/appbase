package migrations

import (
	"fmt"

	"github.com/silvercory/appbase/config"

	"github.com/rs/zerolog"

	_ "github.com/go-sql-driver/mysql"

	"github.com/pressly/goose"
)

func EnsureUpToDate(log zerolog.Logger, conf config.MySQL) error {
	db, err := goose.OpenDBWithDriver(conf.DatabaseType, conf.URI)
	if err != nil {
		return fmt.Errorf("goose: failed to open DB: %w", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Err(err).Msg("goose: failed to close on defer.")
		}
	}()

	if err := goose.Run("up", db, "."); err != nil {
		return fmt.Errorf("goose: run: %w", err)
	}

	return nil
}

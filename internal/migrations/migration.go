package migrations

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/pressly/goose"
)

func Up(db *sql.DB, dir, dialect string) (err error) {
	if _, err = os.Stat(filepath.Base(dir)); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	if err = goose.SetDialect(dialect); err != nil {
		return err
	}

	return goose.Up(db, dir)
}

func Down(db *sql.DB, dir, dialect string) (err error) {
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	if err = goose.SetDialect(dialect); err != nil {
		return err
	}

	return goose.Down(db, dir)

}

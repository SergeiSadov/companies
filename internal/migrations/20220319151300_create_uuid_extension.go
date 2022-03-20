package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateUuidExtension, downCreateUuidExtension)
}

func upCreateUuidExtension(tx *sql.Tx) error {
	if _, err := tx.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`); err != nil {
		return err
	}

	return nil
}

func downCreateUuidExtension(tx *sql.Tx) error {
	if _, err := tx.Exec(`DROP EXTENSION IF NOT EXISTS "uuid-ossp";`); err != nil {
		return err
	}

	return nil
}

package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableIndustries, downCreateTableIndustries)
}

func upCreateTableIndustries(tx *sql.Tx) error {
	if _, err := tx.Exec(`
CREATE TABLE IF NOT EXISTS industries
(
    id            UUID      DEFAULT uuid_generate_v4() PRIMARY KEY,
    name          VARCHAR(255) NOT NULL,
    market_value  INT          NOT NULL,
    co2_footprint VARCHAR(500),
    created       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT idx_industries_unique_name UNIQUE (name)
);
`); err != nil {
		return err
	}

	return nil
}

func downCreateTableIndustries(tx *sql.Tx) error {
	if _, err := tx.Exec(`DROP TABLE industries`); err != nil {
		return err
	}

	return nil
}

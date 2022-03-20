package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableCustomers, downCreateTableCustomers)
}

func upCreateTableCustomers(tx *sql.Tx) error {
	if _, err := tx.Exec(`
CREATE TABLE IF NOT EXISTS companies
(
    id      uuid      DEFAULT UUID_GENERATE_V4() PRIMARY KEY,
    address JSON,
    name    VARCHAR(255),
    industry_id UUID NOT NULL,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_companies_industries FOREIGN KEY (industry_id) REFERENCES industries (id)
);`); err != nil {
		return err
	}

	return nil
}

func downCreateTableCustomers(tx *sql.Tx) error {
	if _, err := tx.Exec(`DROP TABLE companies`); err != nil {
		return err
	}

	return nil
}

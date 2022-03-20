package configs

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"

	"companies/pkg/constants"
)

type Config struct {
	App      App
	Logger   LoggerParams
	Database Database
}

type App struct {
	Environment string `envconfig:"ENV" default:"develop"`
	Port        int    `envconfig:"PORT" required:"true"`
	Name        string `envconfig:"NAME" required:"true"`
}

type LoggerParams struct {
	Level string `envconfig:"LOG_LEVEL" required:"true"`
}

type Database struct {
	Host         string `envconfig:"DB_HOST" required:"true"`
	Port         string `envconfig:"DB_PORT" required:"true"`
	Username     string `envconfig:"DB_USERNAME" required:"true"`
	Password     string `envconfig:"DB_PASSWORD" required:"true"`
	Name         string `envconfig:"DB_NAME" required:"true"`
	Dialect      string `envconfig:"DB_DIALECT" required:"true"`
	MigrationDir string `envconfig:"DB_MIGRATION_DIR" required:"true"`
}

func Setup() (config Config, err error) {
	if err = envconfig.Process("", &config); err != nil {
		return
	}

	return
}

func (c *Database) PrepareDSN() (dsn string) {
	return fmt.Sprintf(constants.DSNTemplate, c.Host, c.Username, c.Password, c.Name, c.Port)
}

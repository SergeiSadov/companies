package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"companies/internal/configs"
	"companies/internal/db"
	"companies/internal/migrations"
)

const (
	migrateActionUp   = "up"
	migrateActionDown = "down"
)

var migrateCmd = &cobra.Command{
	Use: "migrate",
}
var migrateUpCmd = &cobra.Command{
	Use: migrateActionUp,
	Run: RunMigration(migrateActionUp),
}
var migrateDownCmd = &cobra.Command{
	Use: migrateActionDown,
	Run: RunMigration(migrateActionDown),
}

func RunMigration(action string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		cfg, err := configs.Setup()
		if err != nil {
			log.Fatal(err)
		}

		database, err := db.New(cfg.Database.Dialect, cfg.Database.PrepareDSN())
		if err != nil {
			log.Fatal(err)
		}

		conn, err := database.GetUnderlyingConn()
		if err != nil {
			log.Fatal(err)
		}

		switch action {
		case migrateActionUp:
			fmt.Println("running migration up")
			if err = migrations.Up(conn, cfg.Database.MigrationDir, cfg.Database.Dialect); err != nil {
				log.Fatal(err)
			}
		case migrateActionDown:
			fmt.Println("running migration down")
			if err = migrations.Down(conn, cfg.Database.MigrationDir, cfg.Database.Dialect); err != nil {
				log.Fatal(err)
			}
		}
	}
}

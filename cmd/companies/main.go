package main

import (
	"log"
)

func main() {
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)
	rootCmd.AddCommand(migrateCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

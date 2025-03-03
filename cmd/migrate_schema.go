package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"golang-project-layout/migrations/schema"
)

// migrateSchemaCmd represents the schema command in Cobra Command structure
var migrateSchemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "migrate the database schema for the go-project",
	Run:   runMigrateSchemaCmd,
}

// init adds the schema command into the migrate command
func init() {
	migrateCmd.AddCommand(migrateSchemaCmd)
}

// runMigrateDataCmd initialize a new connection and create database schema structure
func runMigrateSchemaCmd(cmd *cobra.Command, args []string) {
	databaseConnection := newDatabaseConnection()
	databaseInstance, err := databaseConnection.Open()
	if err != nil {
		log.Fatal("database error:", err)
	}

	err = schema.Migrate(databaseInstance)
	if err != nil {
		log.Fatal("schema migration error:", err)
	}

	log.Println("schema migration completed")
}

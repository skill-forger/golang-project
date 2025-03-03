package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"golang-project-layout/migrations/data"
)

// migrateDataCmd represents the data command in Cobra Command structure
var migrateDataCmd = &cobra.Command{
	Use:   "data",
	Short: "seed the database data for the go-project",
	Run:   runMigrateDataCmd,
}

// init adds the data command into the migrate command
func init() {
	migrateCmd.AddCommand(migrateDataCmd)
}

// runMigrateDataCmd initialize a new connection and seed data into database
func runMigrateDataCmd(cmd *cobra.Command, args []string) {
	databaseConnection := newDatabaseConnection()
	databaseInstance, err := databaseConnection.Open()
	if err != nil {
		log.Fatal("database error:", err)
	}

	err = data.Migrate(databaseInstance)
	if err != nil {
		log.Fatal("data migration error:", err)
	}

	log.Println("data migration completed")
}

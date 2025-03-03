package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"golang-project-layout/database"
	"golang-project-layout/static"
)

// migrateCmd represents the migrate command in Cobra Command structure
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "go-project migrate command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate called")
	},
}

// init adds the migrate command into the root command
func init() {
	rootCmd.AddCommand(migrateCmd)
}

// newDatabaseConnection returns new database connection
func newDatabaseConnection() database.Connection {
	databaseSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString(static.EnvDbUser),
		viper.GetString(static.EnvDbPassword),
		viper.GetString(static.EnvDbHost),
		viper.GetString(static.EnvDbPort),
		viper.GetString(static.EnvDbName),
	)

	return database.NewConnection(databaseSourceName, nil)
}

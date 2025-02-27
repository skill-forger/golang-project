package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrateSchemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "migrate the database schema for the go-project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrateSchema called")
	},
}

// init adds the schema command into the migrate command
func init() {
	migrateCmd.AddCommand(migrateSchemaCmd)
}

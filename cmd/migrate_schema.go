package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrateSchemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrateSchema called")
	},
}

func init() {
	migrateCmd.AddCommand(migrateSchemaCmd)
}

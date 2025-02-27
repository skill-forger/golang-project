package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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

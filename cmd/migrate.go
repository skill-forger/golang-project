package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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

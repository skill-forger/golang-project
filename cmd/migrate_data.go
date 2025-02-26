package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrateDataCmd = &cobra.Command{
	Use:   "data",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrateData called")
	},
}

func init() {
	migrateCmd.AddCommand(migrateDataCmd)
}

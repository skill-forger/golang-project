package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "go-project",
	Short: "go-project root command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("golang project layout root command called")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.Flags().String("config", "./local.env", "config file")
	viper.SetConfigFile(rootCmd.Flag("config").Value.String())

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	viper.AutomaticEnv()
}

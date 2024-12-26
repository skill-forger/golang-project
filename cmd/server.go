package cmd

import (
	"fmt"
	"os"

	"golang-project-layout/infra"
	"golang-project-layout/servers"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func LoadConfig(path string) *infra.AppConfig {
	viper.AddConfigPath(path)
	viper.SetConfigFile("dev.env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.AutomaticEnv()

	return infra.NewAppConfig()
}

func NewServerCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "server",
		Short: "Start the server",

		Run: func(cmd *cobra.Command, args []string) {
			serverEnv := os.Getenv("SERVER_ENV")
			var app *infra.AppConfig
			if serverEnv == "dev" {
				app = LoadConfig(".")
			}

			servers.InitServer(app)
		},
	}

	return command
}

package cmd

import (
	"fmt"
	"os"

	"golang-project-layout/config"
	"golang-project-layout/server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go run cmd/main",
	Short: "A brief description",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(NewServerCommand())
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.golang-project-layout.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func NewServerCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "server",
		Short: "Start the server",

		Run: func(cmd *cobra.Command, args []string) {
			serverEnv := os.Getenv("SERVER_ENV")
			var app *config.AppConfig
			if serverEnv == "dev" {
				app = LoadConfig(".")
			}

			server.InitServer(app)
		},
	}

	return command
}

func LoadConfig(path string) *config.AppConfig {
	viper.AddConfigPath(path)
	viper.SetConfigFile("dev.env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.AutomaticEnv()

	return server.NewAppConfig()
}

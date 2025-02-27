package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"golang-project-layout/database"
	"golang-project-layout/internal/registry"
	"golang-project-layout/server"
)

// serveCmd represents the serve command in Cobra Command structure
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "initialize and start up the go-project server",
	Run:   runServeCmd,
}

// init adds the serve command into the root command
func init() {
	rootCmd.AddCommand(serveCmd)
}

// runServeCmd executes the core logic of the serve command
func runServeCmd(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"),
	)

	databaseConnection := database.NewConnection(dsn, nil)
	databaseInstance, err := databaseConnection.Open()
	if err != nil {
		log.Fatal("database error:", err)
	}

	err = databaseConnection.Ping()
	if err != nil {
		log.Fatal("database error:", err)
	}

	handlerRegistries := registry.NewHandlerRegistries(databaseInstance)

	serverConfigs := []server.ConfigProvider{
		func(e *echo.Echo) { e.Debug = true },
	}

	serverEngine := server.NewEngine(viper.GetString("SERVER_ADDRESS"), serverConfigs...)

	go func() {
		err = serverEngine.Startup(handlerRegistries...)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("server error:", err)
		}
	}()

	<-c

	err = databaseConnection.Close()
	if err != nil {
		log.Println(err)
	}

	err = serverEngine.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}

	log.Println("go project server gracefully shutdowns")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
}

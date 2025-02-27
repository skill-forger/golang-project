package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang-project-layout/database"
	"golang-project-layout/server"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"

	"golang-project-layout/internal/registry"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "initialize and start up the go-project server",
	Run:   runServeCmd,
}

// init adds the serve command into the root command
func init() {
	rootCmd.AddCommand(serveCmd)
}

func runServeCmd(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	databaseConnection := database.NewConnection("", database.NewDefaultConfig())
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

	serverEngine := server.NewEngine("", serverConfigs...)
	err = serverEngine.Startup(handlerRegistries...)
	if err != nil {
		log.Fatal("server error:", err)
	}

	<-c

	err = databaseConnection.Close()
	if err != nil {
		log.Println(err)
	}

	err = serverEngine.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
}

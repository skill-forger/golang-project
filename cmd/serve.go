package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"

	"golang-project-layout/database"
	"golang-project-layout/server"

	authHdl "golang-project-layout/internal/handler/authentication"
	userRepo "golang-project-layout/internal/repository/user"
	authSvc "golang-project-layout/internal/service/authentication"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "initialize and start up the go-project server",
	Run:   runServeCmd,
}

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

	serverConfigs := []server.ConfigProvider{
		func(e *echo.Echo) { e.Debug = true },
	}

	serverEngine := server.NewEngine("", serverConfigs...)

	authRepo := userRepo.NewRepository(databaseInstance)
	authService := authSvc.NewService(authRepo)
	authHandler := authHdl.NewHandler(authService)

	err = serverEngine.Startup()
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

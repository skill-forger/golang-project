package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang-project-layout/database"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "go-project serve command",
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

	db, err := databaseConnection.Open()
	if err != nil {
		log.Fatal(err)
	}

	<-c

	err = databaseConnection.Close()
	if err != nil {
		log.Println(err)
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
}

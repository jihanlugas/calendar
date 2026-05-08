package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jihanlugas/calendar/config"
	"github.com/jihanlugas/calendar/db"
	"github.com/jihanlugas/calendar/router"
)

func runServer() {
	var err error

	// Initialize global DB connection
	globalDB, closeConn := db.GetConnection()
	db.SetGlobalDB(globalDB)
	defer closeConn()

	r := router.Init()

	if err != nil {
		r.Logger.Fatal(err)
	}

	// Start server
	go func() {
		var err error
		err = r.Start(fmt.Sprintf(":%s", config.Server.Port))
		if err != nil && err != http.ErrServerClosed {
			r.Logger.Fatal("Shutting down the server ", err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = r.Shutdown(ctx)
	if err != nil {
		r.Logger.Fatal(err)
	}
}

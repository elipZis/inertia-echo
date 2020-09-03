package main

import (
	"context"
	"elipzis.com/inertia-echo/repository"
	"elipzis.com/inertia-echo/router"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"os/signal"
	"time"
)

// A runnable type
type Runnable interface {
	Run()
	Shutdown(ctx context.Context)
}

// LOAD "*",8,1
func main() {
	// Connect to the database and defer the closing to the end of the process/main
	db := repository.NewDatabase()
	defer db.Conn.Close()

	// Init Echo
	r := router.NewRouter()
	r.Register(r.Echo.Group(""))

	// RUN
	run(r)
}

// RUN
func run(runners ...Runnable) {
	// Start everything
	for _, runner := range runners {
		runner.Run()
	}

	// Wait for interrupt signal to gracefully shutdown everything after a timeout of 10 seconds
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Try to gracefully shutdown everything
	for _, runner := range runners {
		runner.Shutdown(ctx)
	}

}

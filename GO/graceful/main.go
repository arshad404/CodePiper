package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router
	router := gin.Default()

	// Example route
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// Create a server with the Gin router
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	// Start the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// // Use a channel to listen for the SIGINT and SIGTERM signals
	stop := make(chan os.Signal, 1)
	// // kill (no param) default send syscanll.SIGTERM
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGALRM)

	// // Wait for the termination signal
	<-stop

	log.Println("Received termination signal. Shutting down gracefully...")

	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server with the given context
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	// Will take 2 seconds to close the active connections
	go closingDBConnection()

	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}

	log.Println("Server gracefully stopped.")
}

func closingDBConnection() {
	fmt.Println("DB connection will be closed within 2 seconds")
	time.Sleep(time.Second * 2)
	fmt.Println("DB connection is closed now")
}

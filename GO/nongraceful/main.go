package main

import (
	"context"
	"log"
	"net/http"

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

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}

	// Shutdown the server with the given context
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
}

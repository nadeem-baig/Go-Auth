package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nadeem-baig/go-auth/service/user"
)

// StartServer sets up the HTTP server configurations.
func StartServer() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if not specified
    }

    handler := user.NewHandler()

    server := &http.Server{
        Addr:         fmt.Sprintf(":%s", port),
        Handler:      handler, // This is now correct as config.Handler implements http.Handler
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
        IdleTimeout:  60 * time.Second,
    }

    fmt.Printf("Server starting on port %s\n", port)
    if err := server.ListenAndServe(); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}


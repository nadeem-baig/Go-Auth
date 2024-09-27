package main

import (
	"log"

	"github.com/nadeem-baig/go-auth/cmd/api"
)


func main() {
    // Log startup message
    log.Println("Starting the Go HTTP server...")

    // Start HTTP server
    api.StartServer()
}

package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/nadeem-baig/go-auth/config"
	"github.com/nadeem-baig/go-auth/service/user"
	"github.com/nadeem-baig/go-auth/service/post"
	"github.com/nadeem-baig/go-auth/service/payment"
)

// StartServer sets up the HTTP server configurations.
func StartServer(db *sql.DB) {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	mainHandler := config.NewHandler(db)

	// Initialize service handlers
	userHandler := user.NewHandler(mainHandler)
	postHandler := post.NewHandler(mainHandler)
	paymentHandler := payment.NewHandler(mainHandler)

	// Mount service handlers
	mainHandler.Mux.Handle("/api/v1/users/", http.StripPrefix("/api/v1/users", userHandler))
	mainHandler.Mux.Handle("/api/v1/posts/", http.StripPrefix("/api/v1/posts", postHandler))
	mainHandler.Mux.Handle("/api/v1/payments/", http.StripPrefix("/api/v1/payments", paymentHandler))

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      mainHandler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Printf("Server starting on port %s\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

package user

import (
	"net/http"

	"github.com/nadeem-baig/go-auth/config"
	"github.com/nadeem-baig/go-auth/middleware"
)

// NewHandler initializes a new router with defined routes.
func NewHandler() *config.Handler {
	mux := http.NewServeMux()
	handler := &config.Handler{Mux: mux}

	// Create a sub-router for /api/v1
	apiV1 := http.NewServeMux()

	// Define routes for the sub-router
	apiV1.HandleFunc("GET /", middleware.LoggingMiddleware(HomeHandler))
	apiV1.HandleFunc("GET /login", middleware.LoggingMiddleware(GreetHandler))
	apiV1.HandleFunc("POST /register", middleware.LoggingMiddleware(PostHandler))

	// Mount the sub-router under /api/v1 for all HTTP methods
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiV1))

	return handler
}

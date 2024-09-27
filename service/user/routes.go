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

	// Define routes and their respective handlers
	mux.HandleFunc("/", middleware.LoggingMiddleware(HomeHandler))
	mux.HandleFunc("/greet", middleware.LoggingMiddleware(GreetHandler))
	mux.HandleFunc("/post", middleware.LoggingMiddleware(PostHandler))

	return handler
}


package user

import (
	"net/http"

	"github.com/nadeem-baig/go-auth/config"
	"github.com/nadeem-baig/go-auth/middleware"
)

// NewHandler initializes a new router with defined routes for user service.
func NewHandler(h *config.Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", middleware.LoggingMiddleware(HomeHandler(h)))
	mux.HandleFunc("/login", middleware.LoggingMiddleware(GreetHandler(h)))
	mux.HandleFunc("/register", middleware.LoggingMiddleware(PostHandler(h)))

	return mux
}

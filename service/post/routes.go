package post

import (
	"net/http"

	"github.com/nadeem-baig/go-auth/config"
	"github.com/nadeem-baig/go-auth/middleware"
)

func NewHandler(h *config.Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", middleware.LoggingMiddleware(ListPostsHandler(h)))
	mux.HandleFunc("/create", middleware.LoggingMiddleware(CreatePostHandler(h)))
	// Add more post-related routes

	return mux
}
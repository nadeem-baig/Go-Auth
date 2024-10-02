package user

import (
	"net/http"

	"github.com/nadeem-baig/go-auth/config"
	"github.com/nadeem-baig/go-auth/middleware"
	"github.com/nadeem-baig/go-auth/types"
)

type Handler struct {
	store *types.UserStore
}


// NewHandler initializes a new router with defined routes for user service.
func NewHandler(h *config.Handler) http.Handler {
	mux := http.NewServeMux()
	store := NewStore(h.DB)

	mux.HandleFunc("/", middleware.LoggingMiddleware(HomeHandler(h)))
	mux.HandleFunc("/login", middleware.LoggingMiddleware(LoginHandler(h, store)))
	mux.HandleFunc("/register", middleware.LoggingMiddleware(RegisterHandler(h, store)))
	return mux
}

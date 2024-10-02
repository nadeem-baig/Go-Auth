package product

import (
	"net/http"

	"github.com/nadeem-baig/go-auth/config"
	."github.com/nadeem-baig/go-auth/middleware"
)



// NewHandler initializes a new router with defined routes for user service.
func NewHandler(h *config.Handler) http.Handler {
	mux := http.NewServeMux()
	store := NewStore(h.DB)

	mux.Handle("/list", LoggingMiddleware(MethodHandler(http.MethodGet, GetProducts(h, store))))

	return mux
}

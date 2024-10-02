package product

import (
	"net/http"

	"github.com/nadeem-baig/go-auth/config"
	. "github.com/nadeem-baig/go-auth/middleware"
	"github.com/nadeem-baig/go-auth/service/auth"
	"github.com/nadeem-baig/go-auth/service/user"
)

// NewHandler initializes a new router with defined routes for user service.
func NewHandler(h *config.Handler) http.Handler {
	mux := http.NewServeMux()
	productStore := NewStore(h.DB)
	userStore := user.NewStore(h.DB)

	mux.Handle("/list", LoggingMiddleware(MethodHandler(http.MethodGet, auth.VerifyJWT(GetProducts(h, productStore), userStore))))
	return mux
}

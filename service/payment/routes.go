package payment

import (
	"net/http"

	"github.com/nadeem-baig/go-auth/config"
	"github.com/nadeem-baig/go-auth/middleware"
)

func NewHandler(h *config.Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/process", middleware.LoggingMiddleware(ProcessPaymentHandler(h)))
	mux.HandleFunc("/status", middleware.LoggingMiddleware(PaymentStatusHandler(h)))
	// Add more payment-related routes

	return mux
}
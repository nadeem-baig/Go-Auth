package config

import (
	"net/http"
)



// ServeHTTP allows the handler to be used as an HTTP handler.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Mux.ServeHTTP(w, r)
}

// Response is a struct to define the structure of JSON responses.
type Response struct {
	Message string `json:"message"`
}
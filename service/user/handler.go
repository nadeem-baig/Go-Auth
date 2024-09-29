package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nadeem-baig/go-auth/config"
	"github.com/nadeem-baig/go-auth/utils"
)

// HomeHandler responds with a welcome message.
func HomeHandler(h *config.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.JSONResponse(w, config.Response{Message: "Welcome to the Go HTTP API!"}, http.StatusOK)
	}
}

// GreetHandler responds with a personalized message based on query param.
func GreetHandler(h *config.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}
		utils.JSONResponse(w, config.Response{Message: fmt.Sprintf("Hello, %s!", name)}, http.StatusOK)
	}
}

// PostHandler processes JSON input data and responds.
func PostHandler(h *config.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			utils.JSONResponse(w, config.Response{Message: "Method not allowed"}, http.StatusMethodNotAllowed)
			return
		}

		var data map[string]string
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			utils.JSONResponse(w, config.Response{Message: "Invalid JSON data"}, http.StatusBadRequest)
			return
		}

		utils.JSONResponse(w, config.Response{Message: fmt.Sprintf("Received data: %v", data)}, http.StatusOK)
	}
}
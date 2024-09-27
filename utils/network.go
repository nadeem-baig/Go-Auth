package utils

import (
	"encoding/json"
	"net/http"

	"github.com/nadeem-baig/go-auth/config"
)

// JSONResponse sends a JSON response with the given status code.
func JSONResponse(w http.ResponseWriter, response config.Response, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// Note: The ServeHTTP method is now defined in the config package

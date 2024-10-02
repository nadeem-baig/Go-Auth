package product

import (
	"net/http"

	"github.com/nadeem-baig/go-auth/config"
	"github.com/nadeem-baig/go-auth/utils"
)



// RegisterHandler processes JSON input data and responds.
func GetProducts(h *config.Handler, store ProductStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ps, err := store.GetProducts()
		if err!= nil {
            utils.JSONResponse(w, config.Response{Message: err.Error()}, http.StatusInternalServerError)
            return
        }
		utils.JSONResponse(w, config.Response{Response: ps}, http.StatusOK)


	}
}


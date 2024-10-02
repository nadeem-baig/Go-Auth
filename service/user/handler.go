package user

import (
	"net/http"

	"github.com/nadeem-baig/go-auth/config"
	"github.com/nadeem-baig/go-auth/service/auth"
	"github.com/nadeem-baig/go-auth/types"
	"github.com/nadeem-baig/go-auth/utils"
	"github.com/nadeem-baig/go-auth/utils/logger"
)

// HomeHandler responds with a welcome message.
func HomeHandler(h *config.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.JSONResponse(w, config.Response{Message: "Welcome to the Go HTTP API!"}, http.StatusOK)
	}
}

// RegisterHandler processes JSON input data and responds.
func RegisterHandler(h *config.Handler, store UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload types.RegisterUserPayload

		// Validate and parse JSON request payload
		if err := utils.ParseJson(r, &payload); err != nil {
			utils.JSONResponse(w, config.Response{Message: err.Error()}, http.StatusBadRequest)
			return
		}
		// validate payload
		if err := utils.Validate.Struct(payload); err != nil {
			utils.JSONResponse(w, config.Response{Message: err.Error()}, http.StatusBadRequest)
            return
		}
		// check if user exists
		_, err := store.GetUserByEmail(payload.Email)
		if err == nil {
			utils.JSONResponse(w, config.Response{Message: "User already exists"}, http.StatusConflict)
			return
		}
		hashedPassword,err := auth.HashPassword(payload.Password)
		if err != nil  {
			logger.Errorf("Failed to hash password")
		}
		// create user
		user := types.User{
            FirstName: payload.FirstName,
            LastName:  payload.LastName,
            Email:     payload.Email,
            Password: hashedPassword,
        }
        if err := store.CreateUser(user); err!= nil {
            utils.JSONResponse(w, config.Response{Message: err.Error()}, http.StatusInternalServerError)
            return
        }
        utils.JSONResponse(w, config.Response{Message: "User registered successfully"}, http.StatusCreated)
	}
}

// RegisterHandler processes JSON input data and responds.
func LoginHandler(h *config.Handler, store UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload types.LoginUserPayload

		// Validate and parse JSON request payload
		if err := utils.ParseJson(r, &payload); err != nil {
			utils.JSONResponse(w, config.Response{Message: err.Error()}, http.StatusBadRequest)
			return
		}
		// validate payload
		if err := utils.Validate.Struct(payload); err != nil {
			utils.JSONResponse(w, config.Response{Message: err.Error()}, http.StatusBadRequest)
            return
		}

		u, err := store.GetUserByEmail(payload.Email)
		if err != nil {
			utils.JSONResponse(w, config.Response{Message: "Invalid User details"}, http.StatusBadRequest)
			return
		}
		
		if !auth.ComparePassword(u.Password, payload.Password) {
			utils.JSONResponse(w, config.Response{Message: "Invalid User details"}, http.StatusBadRequest)
            return
		}
		secret := []byte(config.AppConfigs.JWTSecret)
		token,err := auth.CreateJWT(secret,u.ID)
		if err!= nil {
			utils.JSONResponse(w, config.Response{Message:err.Error()}, http.StatusInternalServerError)
            return        
		}
		
        utils.JSONResponse(w, config.Response{Message: token}, http.StatusCreated)
	}
}

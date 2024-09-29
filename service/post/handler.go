package post

import (
	"net/http"

	"github.com/nadeem-baig/go-auth/config"
	"github.com/nadeem-baig/go-auth/utils"
)

func ListPostsHandler(h *config.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implement list posts logic
		utils.JSONResponse(w, config.Response{Message: "List of posts"}, http.StatusOK)
	}
}

func CreatePostHandler(h *config.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implement create post logic
		utils.JSONResponse(w, config.Response{Message: "Post created"}, http.StatusCreated)
	}
}
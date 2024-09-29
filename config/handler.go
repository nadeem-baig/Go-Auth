package config

import (
	"database/sql"
	"net/http"
)

type Handler struct {
	Mux *http.ServeMux
	DB  *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		Mux: http.NewServeMux(),
		DB:  db,
	}
}
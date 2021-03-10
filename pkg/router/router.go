package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/template/pkg/db"
	"github.com/template/pkg/handlers"
	"github.com/template/pkg/middleware"
)

// GetRouter handles routing.
func GetRouter(db *db.DB) *mux.Router {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/users", handlers.GetUsers(db)).Methods(http.MethodGet)

	api.Use(middleware.LoggingMiddleware)

	return api
}

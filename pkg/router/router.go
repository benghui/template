package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/template/pkg/application"
	"github.com/template/pkg/handlers"
	"github.com/template/pkg/middleware"
)

// GetRouter handles routing.
func GetRouter(app *application.Application) *mux.Router {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/users", handlers.GetUsers(app)).Methods(http.MethodGet)

	api.Use(middleware.LoggingMiddleware)

	return api
}

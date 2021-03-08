package handlers

import (
	"net/http"

	"github.com/template/pkg/application"
	"github.com/template/pkg/models"
)

func GetUsers(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := []models.Users{}

		app.DB.Grm.Find(&users)
		respondJSON(w, http.StatusOK, users)
	}
}

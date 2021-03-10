package handlers

import (
	"net/http"

	"github.com/template/pkg/db"
	"github.com/template/pkg/models"
)

func GetUsers(db *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := []models.Users{}

		db.Grm.Find(&users)
		respondJSON(w, http.StatusOK, users)
	}
}

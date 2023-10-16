package util

import (
	"net/http"

	"github.com/autumnleaf-ra/go-movie-api/models"
)

type AppWithCORS struct {
	*models.Application
}

func (app *AppWithCORS) EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

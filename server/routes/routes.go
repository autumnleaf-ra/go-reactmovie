package routes

import (
	"net/http"

	"github.com/autumnleaf-ra/go-movie-api/models"
	"github.com/autumnleaf-ra/go-movie-api/util"
	"github.com/julienschmidt/httprouter"
)

func CreateRouter(app *models.Application) http.Handler {
	router := httprouter.New()

	/* status */
	router.HandlerFunc(http.MethodGet, "/status", StatusHandler(app))
	router.HandlerFunc(http.MethodGet, "/movies/:id", GetOneMovieHandler(app))
	router.HandlerFunc(http.MethodGet, "/movies", GetAllMovieHandler(app))
	router.HandlerFunc(http.MethodGet, "/genres", GetAllMGenreHandler(app))
	router.HandlerFunc(http.MethodGet, "/genres/:genre_id/movies/", GetAllMGenreNameHandler(app))

	router.HandlerFunc(http.MethodPost, "/admin/movies/add", AddMovieHandler(app))
	router.HandlerFunc(http.MethodPost, "/admin/movies/edit", EditMovieHandler(app))
	router.HandlerFunc(http.MethodPost, "/admin/movies/delete", DeleteMovieHandler(app))

	appWithCors := &util.AppWithCORS{Application: app}
	return appWithCors.EnableCORS(router)
}

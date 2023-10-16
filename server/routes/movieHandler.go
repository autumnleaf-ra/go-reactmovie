package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/autumnleaf-ra/go-movie-api/database"
	"github.com/autumnleaf-ra/go-movie-api/entity"
	"github.com/autumnleaf-ra/go-movie-api/models"
	"github.com/autumnleaf-ra/go-movie-api/util"
	"github.com/julienschmidt/httprouter"
)

func GetOneMovie(w http.ResponseWriter, r *http.Request, app *models.Application) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		util.ErrorJSON(w, app, err)
		return
	}

	/* app.Logger.Println("the id is :", id) */
	dbModel := &database.DBModel{DB: app.DB}

	movie, err := dbModel.Get(id)
	if err != nil {
		util.ErrorJSON(w, app, err)
		return
	}

	err = util.WriteJSON(w, app, http.StatusOK, movie, "movie")
	if err != nil {
		util.ErrorJSON(w, app, err)
		return
	}
}

func GetOneMovieHandler(app *models.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		GetOneMovie(w, r, app)
	}
}

func GetAllMovie(w http.ResponseWriter, r *http.Request, app *models.Application) {

	dbModel := &database.DBModel{DB: app.DB}

	movies, err := dbModel.All()
	if err != nil {
		util.ErrorJSON(w, app, err)
		return
	}

	err = util.WriteJSON(w, app, http.StatusOK, movies, "movies")
	if err != nil {
		util.ErrorJSON(w, app, err)
		return
	}
}

func GetAllMovieHandler(app *models.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		GetAllMovie(w, r, app)
	}
}

func GetGenreAll(w http.ResponseWriter, r *http.Request, app *models.Application) {

	dbModel := &database.DBModel{DB: app.DB}

	genres, err := dbModel.GetGenreAll()
	if err != nil {
		util.ErrorJSON(w, app, err)
		return
	}

	err = util.WriteJSON(w, app, http.StatusOK, genres, "genres")
	if err != nil {
		util.ErrorJSON(w, app, err)
		return
	}
}

func GetAllMGenreHandler(app *models.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		GetGenreAll(w, r, app)
	}
}

func GetGenreNameAll(w http.ResponseWriter, r *http.Request, app *models.Application) {
	dbModel := &database.DBModel{DB: app.DB}
	params := httprouter.ParamsFromContext(r.Context())

	genreID, err := strconv.Atoi(params.ByName("genre_id"))
	if err != nil {
		util.ErrorJSON(w, app, err)
		return
	}

	movies, err := dbModel.All(genreID)
	if err != nil {
		util.ErrorJSON(w, app, err)
		return
	}

	err = util.WriteJSON(w, app, http.StatusOK, movies, "movies")
	if err != nil {
		util.ErrorJSON(w, app, err)
		return
	}
}

func GetAllMGenreNameHandler(app *models.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		GetGenreNameAll(w, r, app)
	}
}

func AddMovie(w http.ResponseWriter, r *http.Request, app *models.Application) {
	var payload models.MoviePayload

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		log.Println(err)
		util.ErrorJSON(w, app, err)
		return
	}

	var movie entity.Movie

	movie.ID, _ = strconv.Atoi(payload.ID)
	movie.Title = payload.Title
	movie.Description = payload.Description
	movie.ReleaseDate, _ = time.Parse("2006-01-02", payload.ReleaseDate)
	movie.Year = movie.ReleaseDate.Day()
	movie.Runtime, _ = strconv.Atoi(payload.Runtime)
	movie.Rating, _ = strconv.Atoi(payload.Rating)
	movie.MPAARating = payload.MPAARating
	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()

	/* insert movie query */
	dbModel := &database.DBModel{DB: app.DB}
	dbModel.InsertMovie(movie)
	if err != nil {
		log.Println(err)
		util.ErrorJSON(w, app, err)
		return
	}

	type jsonRes struct {
		OK bool `json:"ok"`
	}

	ok := jsonRes{
		OK: true,
	}

	err = util.WriteJSON(w, app, http.StatusOK, ok, "response")
	if err != nil {
		util.ErrorJSON(w, app, err)
		return
	}
}

func AddMovieHandler(app *models.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		AddMovie(w, r, app)
	}
}

func EditMovie(w http.ResponseWriter, r *http.Request, app *models.Application) {
	var payload models.MoviePayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	dbModel := &database.DBModel{DB: app.DB}

	if err != nil {
		log.Println(err)
		util.ErrorJSON(w, app, err)
		return
	}

	var movie entity.Movie

	id, _ := strconv.Atoi(payload.ID)
	singleMovie, _ := dbModel.Get(id)
	movie = *singleMovie

	movie.ID, _ = strconv.Atoi(payload.ID)
	movie.Title = payload.Title
	movie.Description = payload.Description
	movie.ReleaseDate, _ = time.Parse("2006-01-02", payload.ReleaseDate)
	movie.Year = movie.ReleaseDate.Day()
	movie.Runtime, _ = strconv.Atoi(payload.Runtime)
	movie.Rating, _ = strconv.Atoi(payload.Rating)
	movie.MPAARating = payload.MPAARating
	movie.UpdatedAt = time.Now()

	/* insert movie query */
	dbModel.UpdateMovie(movie)
	if err != nil {
		log.Println(err)
		util.ErrorJSON(w, app, err)
		return
	}

	type jsonRes struct {
		OK bool `json:"ok"`
	}

	ok := jsonRes{
		OK: true,
	}

	err = util.WriteJSON(w, app, http.StatusOK, ok, "response")
	if err != nil {
		util.ErrorJSON(w, app, err)
		return
	}
}

func EditMovieHandler(app *models.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		EditMovie(w, r, app)
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request, app *models.Application) {
	var payload models.MoviePayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	dbModel := &database.DBModel{DB: app.DB}

	if err != nil {
		util.ErrorJSON(w, app, err)
		return
	}

	id, _ := strconv.Atoi(payload.ID)
	err = dbModel.DeleteMovie(id)

	if err != nil {
		util.ErrorJSON(w, app, err)
		return
	}

	type jsonRes struct {
		OK bool `json:"ok"`
	}

	ok := jsonRes{
		OK: true,
	}

	err = util.WriteJSON(w, app, http.StatusOK, ok, "response")
	if err != nil {
		util.ErrorJSON(w, app, err)
		return
	}
}

func DeleteMovieHandler(app *models.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		DeleteMovie(w, r, app)
	}
}

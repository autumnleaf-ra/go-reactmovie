package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/autumnleaf-ra/go-movie-api/constant"
	"github.com/autumnleaf-ra/go-movie-api/models"
)

/* status handler */
func Status(w http.ResponseWriter, r *http.Request, app *models.Application) {
	currentStatus := models.AppStatus{
		Status:      "Online",
		Environment: app.Config.Env,
		Version:     constant.Version,
	}

	res, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func StatusHandler(app *models.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Status(w, r, app)
	}
}

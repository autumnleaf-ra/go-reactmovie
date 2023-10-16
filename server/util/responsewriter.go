package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/autumnleaf-ra/go-movie-api/models"
)

func WriteJSON(w http.ResponseWriter, app *models.Application, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	res, err := json.Marshal(wrapper)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

	return nil
}

func ErrorJSON(w http.ResponseWriter, app *models.Application, err error) {
	type jsonError struct {
		Message string `json:"message"`
	}

	errMessage := jsonError{
		Message: err.Error(),
	}

	WriteJSON(w, app, http.StatusBadRequest, errMessage, "error")
	return
}

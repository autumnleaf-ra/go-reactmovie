package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/autumnleaf-ra/go-movie-api/database"
	"github.com/autumnleaf-ra/go-movie-api/models"
	"github.com/autumnleaf-ra/go-movie-api/routes"
)

func main() {
	var cfg models.Config

	flag.IntVar(&cfg.Port, "port", 4000, "Listen port")
	flag.StringVar(&cfg.Env, "env", "development", "Application environment(development|production)")
	flag.StringVar(&cfg.DB.DSN, "dsn", "postgres://postgres:postgres@localhost/go_reactmovie?sslmode=disable", "Postgres connection config")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := database.OpenDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	app := &models.Application{
		Config: cfg,
		Logger: logger,
		DB:     db,
	}

	fmt.Println("Server is running")

	/* Response Status Server */
	/* http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		currentStatus := models.AppStatus{
			Status:      "Online",
			Environment: cfg.Env,
			Version:     constant.Version,
		}

		res, err := json.MarshalIndent(currentStatus, "", "\t")
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}) */

	serve := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      routes.CreateRouter(app),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting server on port %d", cfg.Port)

	err = serve.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

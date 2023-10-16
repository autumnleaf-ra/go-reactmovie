package models

import (
	"database/sql"
	"log"
)

type Application struct {
	Config Config
	Logger *log.Logger
	DB     *sql.DB
}

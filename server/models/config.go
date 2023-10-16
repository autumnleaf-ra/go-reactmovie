package models

type Config struct {
	Port int
	Env  string
	DB   struct {
		DSN string
	}
}

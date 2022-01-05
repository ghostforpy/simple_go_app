package db

import (
	"os"
	"strings"

	"github.com/go-pg/pg/v10"
)

func NewPostgresDB() *pg.DB {
	conn := pg.Connect(&pg.Options{
		Addr: strings.Join(
			[]string{os.Getenv("POSTGRES_SERVER"), os.Getenv("POSTGRES_PORT")}, ":",
		),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
	})
	return conn
}

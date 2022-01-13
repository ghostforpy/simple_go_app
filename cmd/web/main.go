package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ghostforpy/simple_go_app/internals/db"
	"github.com/ghostforpy/simple_go_app/internals/handlers"
	"github.com/gorilla/mux"
	"github.com/uptrace/bun/extra/bundebug"
)

func main() {

	conn := db.NewPostgresDB()
	conn.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	defer conn.Close()
	ctx := context.Background()
	if _, err := conn.ExecContext(ctx, "SELECT 1"); err != nil {
		panic(err)
	} else {
		log.Println("Connected to DB!")
	}
	r := mux.NewRouter()
	rootRouter := handlers.NewRootHandler()
	rootRouter.RegisterRoutes(conn, ctx, r)

	log.Println("Запуск сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", r)
	log.Fatal(err)
}

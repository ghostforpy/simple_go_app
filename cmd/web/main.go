package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ghostforpy/simple_go_app/internals/db"
	"github.com/ghostforpy/simple_go_app/internals/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	rootRouter := handlers.NewRootHandler()
	rootRouter.RegisterRoutes(r)
	conn := db.NewPostgresDB()
	defer conn.Close()
	ctx := context.Background()

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}
	fmt.Print("all right")
	log.Println("Запуск сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", r)
	log.Fatal(err)
}

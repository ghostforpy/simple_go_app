package handlers

import (
	//"fmt"
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/uptrace/bun"
)

type RootHandler struct {
}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
}

func (h *RootHandler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Привет из Snippetbox"))
}

func (h *RootHandler) RegisterRoutes(conn *bun.DB, ctx context.Context, router *mux.Router) {
	router.HandleFunc("/", h.Home)
	UsersRegisterRoutes(conn, ctx, "/users", router)
}

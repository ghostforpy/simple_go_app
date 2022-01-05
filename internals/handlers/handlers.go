package handlers

import (
	//"fmt"
	"net/http"
	//"strconv"
	"github.com/gorilla/mux"
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

func (h *RootHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", h.Home)
	UsersRegisterRoutes("/users", router)
}

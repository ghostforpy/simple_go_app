package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UsersHandler struct{}

func NewUsersHandler() *UsersHandler {
	return &UsersHandler{}
}

func (h *UsersHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод не дозволен", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Вывод списка всех пользователей"))
}

func (h *UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод не дозволен", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Создание нового пользователя"))
}

func (h *UsersHandler) RetrivieUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Отображение пользователя с ID %d...", id)
}

func (h *UsersHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Изменение пользователя с ID %d...", id)
}
func (h *UsersHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Удаление пользователя с ID %d...", id)
}

func (h *UsersHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", h.ListUsers).Methods("GET")
	router.HandleFunc("/", h.CreateUser).Methods("POST")
	router.HandleFunc("/{id}", h.RetrivieUser).Methods("GET")
	router.HandleFunc("/{id}", h.UpdateUser).Methods("PUT", "PATCH")
	router.HandleFunc("/{id}", h.DeleteUser).Methods("DELETE")
}

func UsersRegisterRoutes(prefix string, router *mux.Router) {
	s := router.PathPrefix(prefix).Subrouter()
	u := NewUsersHandler()
	u.RegisterRoutes(s)
}

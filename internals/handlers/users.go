package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ghostforpy/simple_go_app/internals/crud"
	"github.com/ghostforpy/simple_go_app/internals/models"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/uptrace/bun"
)

type UsersHandler struct {
	conn *bun.DB
	ctx  context.Context
}
type IError struct {
	Field string
	Tag   string
	Value string
}

func NewUsersHandler(conn *bun.DB, ctx context.Context) *UsersHandler {
	return &UsersHandler{conn: conn, ctx: ctx}
}

func GetQueryInt(q url.Values, s string) (int, error) {
	var res int
	var err error
	if len(q[s]) > 0 {
		if res, err = strconv.Atoi(q[s][0]); err == nil {
			return res, nil
		}
	}
	return 0, errors.New("")
}

func (h *UsersHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод не дозволен", http.StatusMethodNotAllowed)
		return
	}
	var users []models.User
	var limit, offset int
	var err error
	queries := r.URL.Query()

	if limit, err = GetQueryInt(queries, "limit"); err != nil || limit <= 0 {
		limit = 0
	}
	if offset, err = GetQueryInt(queries, "offset"); err != nil || offset <= 0 {
		offset = 0
	}
	c := crud.NewUsersCrud(h.conn, h.ctx)
	users, err = c.List(limit, offset)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
		return
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод не дозволен", http.StatusMethodNotAllowed)
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := &models.User{}
	if err = json.Unmarshal(reqBody, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	var Validator = validator.New()
	err = Validator.Struct(user)
	if err != nil {
		var errors []*IError
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
		return
	}
	c := crud.NewUsersCrud(h.conn, h.ctx)
	if user, err := c.Create(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
		return
	}
}

func (h *UsersHandler) RetrivieUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	c := crud.NewUsersCrud(h.conn, h.ctx)
	user, err := c.Retrivie(int64(id))
	if err == nil {
		data, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		}
		return
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}

}

func (h *UsersHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c := crud.NewUsersCrud(h.conn, h.ctx)
	user, err := c.Update(int64(id), reqBody)
	if err == nil {
		user.Password = "" // Don't encode password
		if data, err := json.Marshal(user); err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			return
		}
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func (h *UsersHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	c := crud.NewUsersCrud(h.conn, h.ctx)
	res, err := c.Delete(int64(id))
	if err == nil {
		if res {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *UsersHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", h.ListUsers).Methods("GET")
	router.HandleFunc("/", h.CreateUser).Methods("POST")
	router.HandleFunc("/{id}", h.RetrivieUser).Methods("GET")
	router.HandleFunc("/{id}", h.UpdateUser).Methods("PUT", "PATCH")
	router.HandleFunc("/{id}", h.DeleteUser).Methods("DELETE")
}

func UsersRegisterRoutes(conn *bun.DB, ctx context.Context, prefix string, router *mux.Router) {
	s := router.PathPrefix(prefix).Subrouter()
	u := NewUsersHandler(conn, ctx)
	u.RegisterRoutes(s)
}

package dto

import (
	"time"
)

type UserCreateForAdmin struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" validate:"required,min=2,max=100"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password,omitempty"`
	IsSuperuser bool   `json:"is_superuser"`
	IsActive    bool   `json:"is_active"`
}

type UserCreateForUser struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty"`
}

type User struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" validate:"required,min=2,max=100"`
	Email       string    `json:"email" validate:"required,email"`
	CreatedAt   time.Time `json:"created_at"`
	IsSuperuser bool      `json:"is_superuser"`
	IsActive    bool      `json:"is_active"`
}

package models

import (
	"time"

	"github.com/ghostforpy/simple_go_app/internals/dto"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID          int64     `bun:"id,pk,autoincrement" json:"id"`
	Name        string    `bun:"name,notnull" json:"name" validate:"required,min=2,max=100"`
	Email       string    `bun:",unique,notnull" json:"email" validate:"required,email"`
	Password    string    `bun:",notnull" json:"password,omitempty"`
	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	IsSuperuser bool      `bun:",notnull,default:false" json:"is_superuser"`
	IsActive    bool      `bun:",notnull,default:false" json:"is_active"`
}

func UserToDTO(user *User) dto.User {
	return dto.User{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		CreatedAt:   user.CreatedAt,
		IsSuperuser: user.IsSuperuser,
		IsActive:    user.IsActive,
	}
}

package crud

import (
	//"database/sql"
	//"fmt"
	//"os"
	"context"
	"encoding/json"
	"fmt"

	"github.com/ghostforpy/simple_go_app/internals/dto"
	"github.com/ghostforpy/simple_go_app/internals/models"

	"github.com/uptrace/bun"
	//"github.com/uptrace/bun/dialect/pgdialect"
	//"github.com/uptrace/bun/driver/pgdriver"
)

type UserRepo interface {
	List(limit, offset int) ([]dto.User, error)
	Create(user *models.User) (*models.User, error)
	Retrivie(id int64) (*models.User, error)
	Update(id int64, reqBody []byte) (*models.User, error)
	Delete(id int64) (bool, error)
}

type UsersCrud struct {
	conn *bun.DB
	ctx  context.Context
}

func NewUsersCrud(conn *bun.DB, ctx context.Context) *UsersCrud {
	return &UsersCrud{conn: conn, ctx: ctx}
}
func (c *UsersCrud) List(limit, offset int) ([]dto.User, error) {
	var users []models.User
	fmt.Println(limit, offset)
	q := c.conn.NewSelect().Model(&users).ExcludeColumn("password").OrderExpr("id")
	if limit > 0 {
		q = q.Limit(limit)
	}
	if offset > 0 {
		q = q.Offset(offset)
	}
	err := q.Scan(c.ctx)
	if err == nil {
		var u []dto.User
		for _, i := range users {
			u = append(u, models.UserToDTO(&i))
		}
		return u, nil
	} else {
		return nil, err
	}
}

func (c *UsersCrud) Create(user *models.User) (*models.User, error) {
	_, err := c.conn.NewInsert().Model(user).Returning("*").Exec(c.ctx)
	return user, err
}

func (c *UsersCrud) Retrivie(id int64) (*models.User, error) {
	user := &models.User{ID: id}
	err := c.conn.NewSelect().Model(user).WherePK().ExcludeColumn("password").Scan(c.ctx)
	return user, err
}

func (c *UsersCrud) Update(id int64, reqBody []byte) (*models.User, error) {
	user := &models.User{ID: int64(id)}
	err := c.conn.NewSelect().Model(user).WherePK().
		ExcludeColumn("created_at").Scan(c.ctx)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(reqBody, &user)
	_, err = c.conn.NewUpdate().
		Model(user).
		WherePK().
		ExcludeColumn("created_at").
		Returning("*").
		Exec(c.ctx)
	return user, err
}

func (c *UsersCrud) Delete(id int64) (bool, error) {
	user := &models.User{ID: id}
	res, err := c.conn.NewDelete().Model(user).WherePK().Returning("*").Exec(c.ctx)
	if err == nil {
		if c, _ := res.RowsAffected(); c > 0 {
			return true, nil
		} else {
			return false, nil
		}
	}
	return false, err
}

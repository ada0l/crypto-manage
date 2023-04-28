package database

import (
	"context"
	"crypto-manager/model"
	"github.com/jmoiron/sqlx"
	"gopkg.in/telebot.v3"
)
import _ "embed"

//go:embed sql/user/get_by_id.sql
var userGetByIdQuery string

func UserGetById(ctx context.Context, db *sqlx.DB, tgUser *telebot.User) (user model.User, err error) {
	err = db.GetContext(ctx, &user, userGetByIdQuery, tgUser.ID)
	return user, err
}

//go:embed sql/user/create.sql
var userCreateQuery string

func UserCreate(ctx context.Context, db *sqlx.DB, user model.User) (err error) {
	_, err = db.NamedExecContext(ctx, userCreateQuery, user)
	return err
}

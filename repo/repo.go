package repo

import (
	"context"
	"database/sql"

	"github.com/mohashari/kit-user/domain"
)

type UserRepo interface {
	GetUser(ctx context.Context) (interface{}, error)
	CreateUser(ctx context.Context, user domain.User) (interface{}, error)
}

type userRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{
		Db: db,
	}
}

func (r *userRepo) GetUser(ctx context.Context) (result interface{}, err error) {
	return "", nil
}

func (r *userRepo) CreateUser(ctx context.Context, user domain.User) (result interface{}, err error) {
	return "", nil
}

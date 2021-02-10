package service

import (
	"context"

	"github.com/mohashari/kit-user/domain"
	"github.com/mohashari/kit-user/repo"
)

type UserService interface {
	GetUser(ctx context.Context) (interface{}, error)
	CreateUser(ctx context.Context, user domain.User) (interface{}, error)
}

type userService struct {
	userRepo repo.UserRepo
}

func NewUserService(repo repo.UserRepo) UserService {
	return &userService{
		userRepo: repo,
	}
}

func (s *userService) GetUser(ctx context.Context) (interface{}, error) {
	return "", nil
}

func (s *userService) CreateUser(ctx context.Context, user domain.User) (interface{}, error) {
	return "", nil
}

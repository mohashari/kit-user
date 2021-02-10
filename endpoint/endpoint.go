package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mohashari/kit-user/domain"
	"github.com/mohashari/kit-user/service"
)

type UserEndpoint struct {
	CreateUserEndpoint endpoint.Endpoint
	GetUserEndpoint    endpoint.Endpoint
}

func NewUserEndpoint(service service.UserService) UserEndpoint {
	return UserEndpoint{
		CreateUserEndpoint: makeCreateUserEndpoint(service),
		GetUserEndpoint:    makeGetUserEndpoint(service),
	}
}

func makeGetUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.GetUser(ctx)
	}
}

func makeCreateUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(domain.User)
		return svc.CreateUser(ctx, req)
	}
}

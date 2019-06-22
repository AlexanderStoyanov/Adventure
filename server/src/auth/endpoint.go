package register

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	userpkg "github.com/AlexanderStoyanov/Adventure/server/src/user"
)

// RegisterUserRequest holds the request parameters for the RegisterUser method
type RegisterUserRequest struct {
	User userpkg.User
}

// RegisterUserResponse holds the response values for the RegisterUser method
type RegisterUserResponse struct {
	Err error
}

func makeRegisterUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterUserRequest)
		err := s.RegisterUser(ctx, req.User)
		return RegisterUserResponse{Err: err}, nil
	}
}

// LoginUserRequest holds the request parameters for the LoginUser method
type LoginUserRequest struct {
	Username string
	Password string
}

// LoginUserResponse holds the response values for the LoginUser method
type LoginUserResponse struct {
	Err error
}

func makeLoginUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginUserRequest)
		err := s.LoginUser(ctx, req.Username, req.Password)
		return LoginUserResponse{Err: err}, nil
	}
}
package auth

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/go-kit/kit/endpoint"
)

// RegisterUserRequest holds the request parameters for the RegisterUser method
type RegisterUserRequest struct {
	User *auth.UserToCreate
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
	Email    string
	Password string
}

// LoginUserResponse holds the response values for the LoginUser method
type LoginUserResponse struct {
	User *auth.UserRecord
	Err  error
}

func makeLoginUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginUserRequest)
		user, err := s.LoginUser(ctx, req.Email, req.Password)
		return LoginUserResponse{User: user, Err: err}, nil
	}
}

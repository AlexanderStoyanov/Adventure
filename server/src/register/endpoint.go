package register

import (
	"context"
	userpkg "user"

	"github.com/go-kit/kit/endpoint"
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

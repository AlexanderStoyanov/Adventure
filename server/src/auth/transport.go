package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"firebase.google.com/go/auth"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// MakeHandler returns a handler for the register service
func MakeHandler(as Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	registerUserHandler := kithttp.NewServer(
		makeRegisterUserEndpoint(as),
		decodeRegisterUserRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/auth/register", registerUserHandler).Methods("POST")

	return r
}

func decodeRegisterUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	params := (&auth.UserToCreate{}).
		Email(body.Email).
		EmailVerified(false).
		Password(body.Password).
		DisplayName(body.Username).
		Disabled(false)

	return RegisterUserRequest{
		params,
	}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case ErrInsertInRepository:
		w.WriteHeader(http.StatusBadRequest)
	case ErrQueryRepository:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

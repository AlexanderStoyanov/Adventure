package event

import (
	"context"
	"encoding/json"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// MakeHandler returns a handler for the register service
func MakeHandler(es Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	createEventHandler := kithttp.NewServer(
		makeCreateEventEndpoint(es),
		decodeCreateEventRequest,
		encodeResponse,
		opts...,
	)

	getEventListHandler := kithttp.NewServer(
		makeGetEventListEndpoint(es),
		decodeGetEventListRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/event/create", createEventHandler).Methods("POST")
	r.Handle("/event/getall", getEventListHandler).Methods("GET")

	return r
}

func decodeCreateEventRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	params := New(body.Name)

	return CreateEventRequest{
		params,
	}, nil
}

func decodeGetEventListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return GetEventListRequest{}, nil
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

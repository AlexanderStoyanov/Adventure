package event

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// CreateEventRequest holds the request parameters for the CreateEvent method
type CreateEventRequest struct {
	Event *Event
}

// CreateEventResponse holds the response values for the CreateEvent method
type CreateEventResponse struct {
	Err error
}

func makeCreateEventEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateEventRequest)
		err := s.CreateEvent(ctx, req.Event)
		return CreateEventResponse{Err: err}, nil
	}
}

// GetEventListRequest holds the request parameters for the GetEventByID method
type GetEventListRequest struct {
}

// GetEventListResponse holds the response values for the GetEventByID method
type GetEventListResponse struct {
	Events []Event `json:"events,omitempty"`
	Err    error   `json:"error,omitempty"`
}

func makeGetEventListEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(GetEventListRequest)
		events, err := s.GetEventList(ctx)
		return GetEventListResponse{Events: events, Err: err}, nil
	}
}

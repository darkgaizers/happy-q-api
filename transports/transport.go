package transports

import (
	"context"
	"encoding/json"
	"happy-q-api/services"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func MakeUppercaseEndpoint(svc services.StringServiceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return UppercaseResponse{v, err.Error()}, nil
		}
		return UppercaseResponse{v, ""}, nil
	}
}

func MakeCountEndpoint(svc services.StringServiceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(CountRequest)
		v := svc.Count(req.S)
		return CountResponse{v}, nil
	}
}

func DecodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type UppercaseRequest struct {
	S string `json:"s"`
}

type UppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type CountRequest struct {
	S string `json:"s"`
}

type CountResponse struct {
	V int `json:"v"`
}

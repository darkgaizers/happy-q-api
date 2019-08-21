package transports

import (
	"context"
	"encoding/json"
	"happy-q-api/interfaces"
	"happy-q-api/models"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func MakePersonGetEndpoint(qs interfaces.PersonServiceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(PersonGetRequest)
		p, err := qs.Get(req.UID)
		if err != nil {
			return PersonAddResponse{Err: err.Error()}, nil
		}
		person := PersonResponse{

			UID:      p.ID,
			Name:     p.Name,
			UserType: p.Type,
		}
		return PersonGetResponse{person, ""}, nil
	}
}

func MakePersonAddEndpoint(qs interfaces.PersonServiceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(PersonAddRequest)

		if err := qs.Add(&models.Person{Name: req.Name}); err != nil {
			return PersonAddResponse{err.Error()}, nil
		}
		return PersonAddResponse{}, nil
	}
}
func MakePersonUpdateEndpoint(qs interfaces.PersonServiceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(PersonUpdateRequest)

		if err := qs.Add(&models.Person{Name: req.Name}); err != nil {
			return PersonUpdateResponse{err.Error()}, nil
		}
		return PersonUpdateResponse{}, nil
	}
}
func DecodePersonAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request PersonAddRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodePersonGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request PersonGetRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
func DecodePersonUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request PersonUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

type PersonAddRequest struct {
	Name string `json:"name"`
}

type PersonAddResponse struct {
	Err string `json:"err,omitempty"`
}
type PersonGetRequest struct {
	UID string `json:"uid"`
}
type PersonResponse struct {
	UID      string `json:"uid"`
	Name     string `json:"name"`
	UserType string `json:"user_type"`
}
type PersonGetResponse struct {
	PersonResponse
	Err string `json:"err,omitempty"`
}
type PersonUpdateRequest struct {
	Name string `json:"name"`
}
type PersonUpdateResponse struct {
	Err string `json:"err,omitempty"`
}

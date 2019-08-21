package transports

import (
	"context"
	"encoding/json"
	"happy-q-api/interfaces"
	"happy-q-api/models"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func MakeQueuePopEndpoint(qs interfaces.QueueServiceInterface, ps interfaces.PersonServiceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(QueuePopRequest)
		v := qs.Pop(&models.Service{ID: req.SID})
		p, err := ps.Get(v.UserID)
		if err != nil {
			if err.Error() != "not implemented" {
				return QueuePopResponse{Err: err.Error()}, nil
			}

		}
		person := PersonResponse{

			UID:      p.ID,
			Name:     p.Name,
			UserType: p.Type,
		}
		return QueuePopResponse{v.No, person, ""}, nil
	}
}

func MakeQueuePushEndpoint(qs interfaces.QueueServiceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(QueuePushRequest)

		v, err := qs.Push(&models.Service{ID: req.SID}, &models.Person{ID: req.UID})
		if err != nil {
			return QueuePushResponse{v.CurrentQueue, err.Error()}, nil
		}
		return QueuePushResponse{v.CurrentQueue, ""}, nil
	}
}
func MakeQueueViewEndpoint(qs interfaces.QueueServiceInterface, ps interfaces.PersonServiceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(QueueViewRequest)
		v, err := qs.View(&models.Service{ID: req.SID})
		if err != nil {
			return QueueViewResponse{Err: err.Error()}, nil
		}
		data := []QueueViewData{}
		/* 		type QueueData struct{
			ID string
			ServiceID string
			UserID string
			CompletedFlag bool
			CreatedDate time.Time
		} */
		for _, d := range v.Data {

			data = append(data, QueueViewData{
				No:  d.No,
				UID: d.UserID,
			})
		}
		return QueueViewResponse{SID: v.ServiceID, TotalQueue: v.TotalQueue, Data: data}, nil
	}
}
func DecodeQueuePushRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request QueuePushRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeQueuePopRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request QueuePopRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
func DecodeQueueViewRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request QueueViewRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

type QueuePushRequest struct {
	SID  string `json:"sid"`
	SSID string `json:"ssid"`
	UID  string `json:"uid"`
}

type QueuePushResponse struct {
	No  int    `json:"no"`
	Err string `json:"err,omitempty"`
}
type QueuePopRequest struct {
	SID string `json:"sid"`
}
type QueuePopResponse struct {
	No int `json:"no"`
	PersonResponse
	Err string `json:"err,omitempty"`
}
type QueueViewRequest struct {
	SID string `json:"sid"`
}
type QueueViewResponse struct {
	SID        string          `json:"sid"`
	TotalQueue int             `json:"total_queue"`
	Data       []QueueViewData `json:"data"`
	Err        string          `json:"err,omitempty"`
}
type QueueViewData struct {
	No   int    `json:"no"`
	UID  string `json:"uid"`
	Name string `json:"user_name"`
}

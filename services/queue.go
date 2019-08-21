package services

import (
	"errors"
	"happy-q-api/models"
)

// StringService provides operations on strings.

/* type QueueServiceInterface interface {
	Push(*models.Service,*models.Person) (*models.QueueResult, error)
	Pop() *models.Person
	View() (*models.QueueView, error)
} */
type QueueService struct{}

func (QueueService) Push(*models.Service, *models.Person) (*models.QueueResult, error) {
	return &models.QueueResult{
		CurrentQueue: 1,
	}, errors.New("not implemented")
}

func (QueueService) Pop(*models.Service) *models.QueueMetadata {
	return &models.QueueMetadata{
		No:        1,
		UserID:    "1",
		ServiceID: "1",
	}
}
func (QueueService) View(*models.Service) (*models.QueueView, error) {
	return &models.QueueView{
		ServiceID:  "1",
		TotalQueue: 0,
		Data:       []models.QueueMetadata{},
	}, errors.New("not implemented")
}

// ErrEmpty is returned when an input string is empty.
/* var popError = errors.New("pop error")
var pushError = errors.New("push error") */

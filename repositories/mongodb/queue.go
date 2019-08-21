package mongodb

import (
	"context"
	"errors"
	"happy-q-api/interfaces"
	"happy-q-api/models"
)
/* type QueueRepository interface {
	Push(*models.Service,*models.Person) (*models.QueueResult, error)
	Pop() *models.Person
	View() (*models.QueueView, error)
} */
type mongoDBQueueRepository struct {
	Conn string
}

func (r *mongoDBQueueRepository) Push(ctx context.Context,service *model.Service, person *models.Person) (*models.QueueResult, error) {
	return nil, errors.New("not implemented")
}
func (r *mongoDBQueueRepository) Pop(ctx context.Context,service *model.Service) (*models.Person, error) {
	return nil, errors.New("not implemented")
}
func (r *mongoDBQueueRepository) View(ctx context.Context,service *model.Service) (*models.QueueView, error) {
	return nil, errors.New("not implemented")
}
func NewMongoDBQueueRepository(conn string) interfaces.QueueRepository {
	return &mongoDBQueueRepository{conn}
}

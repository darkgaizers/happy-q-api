package interfaces
import (
	"context"
	"happy-q-api/models"
)
type QueueServiceInterface interface {
	Push(*models.Service, *models.Person) (*models.QueueResult, error)
	Pop(*models.Service) *models.QueueMetadata
	View(*models.Service) (*models.QueueView, error)
}
type QueueRepository interface {
	Push(ctx context.Context,service *models.Service, person *models.Person) (*models.QueueResult, error)
	Pop(ctx context.Context,service *models.Service) (*models.QueueMetadata, error)
	View(ctx context.Context,service *models.Service) (*models.QueueView, error)
}
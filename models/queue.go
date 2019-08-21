package models
import (
	"time"
)
/* type QueueRepository interface {
	Push(*models.Service,*models.Person) (*models.QueueResult, error)
	Pop() *models.Person
	View() (*models.QueueView, error)
} */
type QueueMetadata struct {
	ServiceID string
	UserID string
	No int
}
type QueueData struct{
	ID string
	QueueMetadata
	
	CompletedFlag bool
	CreatedDate time.Time
}
type QueueResult struct{
	CurrentQueue int
}
type QueueView struct{
	ServiceID string
	QueueResult
	TotalQueue int
	Data []QueueMetadata
}
/* type QueueRepository struct{
	Push(ctx context.Context, employeeID string) (*models.Employee, error)
} */
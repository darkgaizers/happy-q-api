package interfaces
import (
	"context"
	"happy-q-api/models"
)
type PersonServiceInterface interface {
	Get(ID string) (*models.Person,error)
	Add(*models.Person) error
	Update(*models.PersonUpdate) error
}
type PersonRepository interface {
	Get(ctx context.Context ,ID string) (*models.Person, error)
	Add(ctx context.Context,p *models.Person) error
	Update(ctx context.Context,pu *models.PersonUpdate) error
}
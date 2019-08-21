package mongodb

import (
	"context"
	"errors"
	"happy-q-api/interfaces"
	"happy-q-api/models"
)
/* 
	Get(ctx context.Context ,ID string) (*models.Person, error)
	Add(ctx context.Context,p *models.Person) error
	Update(ctx context.Context,pu *models.PersonUpdate) error
*/
type mongoDBPersonRepository struct {
	Conn string
}

func (r *mongoDBPersonRepository) Get(ctx context.Context,ID string) (*models.Person, error) {
	return nil, errors.New("not implemented")
}
func (r *mongoDBPersonRepository) Add(ctx context.Context,p *models.Person) ( error) {
	return  errors.New("not implemented")
}
func (r *mongoDBPersonRepository) Update(ctx context.Context,p *models.PersonUpdate) ( error) {
	return  errors.New("not implemented")
}
func NewMongoDBPersonRepository(conn string) interfaces.PersonRepository {
	return &mongoDBPersonRepository{conn}
}

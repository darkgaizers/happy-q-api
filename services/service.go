package services

import (
	"context"
	"errors"
	"happy-q-api/interfaces"
	"happy-q-api/models"
)

// StringService provides operations on strings.

/* type QueueServiceInterface interface {
	Push(*models.Service,*models.Person) (*models.QueueResult, error)
	Pop() *models.Person
	View() (*models.QueueView, error)
} */
/* Get(ID string) (*models.Person,error)
Add(*models.Person) error
Update(*models.PersonUpdate) error */
type personService struct {
	repository interfaces.PersonRepository
}

func (p *personService) Get(ID string) (*models.Person, error) {
	res, err := p.repository.Get(context.Background(), ID)
	if err != nil {
		return nil, err
	}
	return res, nil

	/* 	return &models.Person{
		ID:   "1",
		Name: "Test",
		Type: "MockupType",
	}, errors.New("not implemented") */
}

func (p *personService) Add(*models.Person) error {
	return errors.New("not implemented")
}
func (p *personService) Update(*models.PersonUpdate) error {
	return errors.New("not implemented")
}
func NewPersonService(er interfaces.PersonRepository) interfaces.PersonServiceInterface {
	return &personService{er}
}

// ErrEmpty is returned when an input string is empty.
/* var popError = errors.New("pop error")
var pushError = errors.New("push error") */

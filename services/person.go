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
/* Get(ID string) (*models.Person,error)
Add(*models.Person) error
Update(*models.PersonUpdate) error */
type PersonService struct{}

func (PersonService) Get(ID string) (*models.Person, error) {
	return &models.Person{
		ID:   "1",
		Name: "Test",
		Type: "MockupType",
	}, errors.New("not implemented")
}

func (PersonService) Add(*models.Person) error {
	return errors.New("not implemented")
}
func (PersonService) Update(*models.PersonUpdate) error {
	return errors.New("not implemented")
}

// ErrEmpty is returned when an input string is empty.
/* var popError = errors.New("pop error")
var pushError = errors.New("push error") */

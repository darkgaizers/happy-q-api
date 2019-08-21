package services

import (
	"context"
	"errors"
	"happy-q-api/interfaces"
	"happy-q-api/models"
)

type personService struct {
	repository interfaces.PersonRepository
}

func (p *personService) Get(ID string) (*models.Person, error) {
	res, err := p.repository.Get(context.Background(), ID)
	if err != nil {
		return nil, err
	}
	return res, nil
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

package peopleService

import (
	models "people-api/internal/app/Models"
	peopleRepo "people-api/internal/app/Repositories/People"
	peopleDto "people-api/internal/app/dtos/People"

	"go.uber.org/zap"
)

type PeopleService interface {
	CreatePerson(req *peopleDto.PeopleCreateRquest) (*models.Person, error)
	GetPerson(id *uint) (*models.Person, error)
}

type PeopleServiceImpl struct {
	PeopleRepo peopleRepo.PeopleRepo
	Logger     *zap.Logger
}

func NewPeopleServiceImpl(logger *zap.Logger) PeopleService {
	return &PeopleServiceImpl{Logger: logger}
}

func (ps *PeopleServiceImpl) GetPerson(id *uint) (*models.Person, error) {

	person, err := ps.PeopleRepo.GetPerson(id)

	if err != nil {
		return nil, err
	}

	return person, nil
}

func (ps *PeopleServiceImpl) CreatePerson(req *peopleDto.PeopleCreateRquest) (*models.Person, error) {

	//Create person model
	person := &models.Person{
		Name:     req.Name,
		Email:    req.Email,
		LastName: req.LastName,
		Age:      req.Age,
	}

	person, err := ps.PeopleRepo.CreatePerson(person)

	if err != nil {
		return nil, err
	}

	return person, nil
}

package peopleService

import (
	"os"
	models "people-api/internal/app/Models"
	peopleRepo "people-api/internal/app/Repositories/People"
	peopleDto "people-api/internal/app/dtos/People"
	"strconv"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type PeopleService interface {
	CreatePerson(req *peopleDto.PeopleCreateRquest) (*models.Person, error)
	GetPerson(id *uint) (*models.Person, error)
	GetPersonByEmail(email *string) (*models.Person, error)
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

func (ps *PeopleServiceImpl) GetPersonByEmail(email *string) (*models.Person, error) {

	person, err := ps.PeopleRepo.GetPersonByEmail(email)

	if err != nil {
		return nil, err
	}

	return person, nil
}

func (ps *PeopleServiceImpl) CreatePerson(req *peopleDto.PeopleCreateRquest) (*models.Person, error) {

	cost, err := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	if err != nil {
		return nil, err
	}

	passwordCrypt, err := bcrypt.GenerateFromPassword([]byte(*req.Password), cost)
	if err != nil {
		return nil, err
	}

	passWord := string(passwordCrypt)
	amount := 0.0

	//Create person model
	person := &models.Person{
		Name:     req.Name,
		Email:    req.Email,
		LastName: req.LastName,
		Age:      req.Age,
		Password: &passWord,
		Amount:   &amount,
	}

	personCreated, err := ps.PeopleRepo.CreatePerson(person)
	if err != nil {
		return nil, err
	}

	return personCreated, nil
}

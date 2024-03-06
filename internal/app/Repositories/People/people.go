package peopleRepo

import (
	models "people-api/internal/app/Models"

	"gorm.io/gorm"
)

type PeopleRepo interface {
	CreatePerson(person *models.Person) (*models.Person, error)
	GetPerson(id *uint) (*models.Person, error)
}

type PeopleRepoImpl struct {
	DB *gorm.DB
}

func NewPeopleRepoImpl(db *gorm.DB) PeopleRepo {
	return &PeopleRepoImpl{
		DB: db,
	}
}

func (pr *PeopleRepoImpl) GetPerson(id *uint) (*models.Person, error) {
	person := &models.Person{}

	tx := pr.DB.Where("id = ?", *id).Find(&person)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return person, nil
}

func (pr *PeopleRepoImpl) CreatePerson(person *models.Person) (*models.Person, error) {
	tx := pr.DB.Create(person)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return person, nil
}

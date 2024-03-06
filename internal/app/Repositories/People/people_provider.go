package peopleRepo

import "gorm.io/gorm"

func InitPeopleRepo(db *gorm.DB) PeopleRepo {
	return &PeopleRepoImpl{DB: db}
}

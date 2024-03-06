package peopleService

import (
	peopleRepo "people-api/internal/app/Repositories/People"
	db "people-api/internal/app/database"

	"go.uber.org/zap"
)

func InitPeopleService(logger *zap.Logger) PeopleService {
	db, _ := db.GetInstance()
	peopleRep := peopleRepo.InitPeopleRepo(db)
	return &PeopleServiceImpl{PeopleRepo: peopleRep, Logger: logger}
}

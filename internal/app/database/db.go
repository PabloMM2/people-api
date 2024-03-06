package db

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}

var DbCon *gorm.DB

func GetInstance() (*gorm.DB, error) {
	if DbCon == nil {
		lock.Lock()
		defer lock.Unlock()
		if DbCon == nil {
			connection, err := initDb()
			if err != nil {
				return nil, err
			}

			DbCon = connection
		}
	}
	return DbCon, nil
}

func initDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("HOST_DB"),
		os.Getenv("USER_DB"),
		os.Getenv("PSWD_DB"),
		os.Getenv("NAME_DB"),
		os.Getenv("PORT_DB"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}

package database

import (
	"main/environment"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {

	var err error

	DB, err = gorm.Open(postgres.Open(environment.Env.DbConnectionString))

	if err != nil {
		panic("falied to connect db. \n\nError: " + err.Error() + "\n")
	}
}

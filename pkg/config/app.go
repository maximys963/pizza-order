package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var postgresDB *gorm.DB

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func Connect(config DBConfig) {
	dbConnectionURL := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName)

	db, err := gorm.Open(postgres.Open(dbConnectionURL), &gorm.Config{
		SkipDefaultTransaction: true,
		QueryFields:            true,
	})

	if err != nil {
		log.Panic(err)
	}

	_, dbErr := db.DB()
	if dbErr != nil {
		log.Panic(dbErr)
	}

	postgresDB = db
}

func GetDatabase() *gorm.DB {
	return postgresDB
}

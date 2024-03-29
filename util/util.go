package util

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func GetEnvOrFail(envName string) string {
	var envValue = os.Getenv(envName)
	if envValue == "" {
		panic("env: " + envName + " must be not empty")
	}

	return envValue
}

func CreateDBConnection(
	dbConnectionURL string,
) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbConnectionURL), &gorm.Config{
		SkipDefaultTransaction: true,
		QueryFields:            true,
	})
	if err != nil {
		log.Panic(err)
	}

	_, dbErr := db.DB()
	if dbErr != nil {
		return nil, err
	}

	return db, err
}

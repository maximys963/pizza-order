package models

import (
	"github.com/maximys963/pizza-order/pkg/config"
	"github.com/sirupsen/logrus"
)

type Film struct {
	Name   string
	Actors string
	Year   int
	Format string
}

func GetAllFilms() []Film {
	var db = config.GetDatabase()
	var films []Film

	result := db.Find(&films)

	if result.Error != nil {
		logrus.Error(result.Error)
	}

	return films
}

func AddFilm(film Film) Film {
	var db = config.GetDatabase()

	result := db.Create(film)

	if result.Error != nil {
		logrus.Error(result.Error)
	}

	return film
}

func DeleteFilm(filmId string) Film {
	var db = config.GetDatabase()

	var film Film

	result := db.Where("ID=?", filmId).Delete(film)

	if result.Error != nil {
		logrus.Error(result.Error)
	}

	return film
}

func ImportFilm() {}
func UpdateFilm() {}

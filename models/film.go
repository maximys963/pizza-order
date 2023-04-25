package models

import (
	"github.com/maximys963/pizza-order/pkg/config"
	"github.com/sirupsen/logrus"
)

type Film struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Actors string `json:"actors"`
	Year   int    `json:"year"`
	Format string `json:"format"`
}

func GetFilm(filmId string) Film {
	var db = config.GetDatabase()

	var film Film

	result := db.Where("ID=?", filmId).Find(&film)

	if result.Error != nil {
		logrus.Error(result.Error)
	}

	return film
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

func UpdateFilm(filmId string, film Film) Film {
	var db = config.GetDatabase()

	result := db.Where("ID=?", filmId).Save(film)

	if result.Error != nil {
		logrus.Error(result.Error)
	}

	return film
}

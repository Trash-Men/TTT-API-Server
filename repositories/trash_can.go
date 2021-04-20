package repositories

import (
	"github.com/Trash-Men/api-server/models"
)

type TrashCanRepository struct{}

func (_ TrashCanRepository) Create(photoUrl string, latitude, longitude float64, area string) error {
	_, error := dbClient.Model(&models.TrashCan{
		Photo_url: photoUrl,
		Latitude:  latitude,
		Longitude: longitude,
		Area:      area,
	}).Insert()

	if error != nil {
		return error
	}

	return nil
}

func (_ TrashCanRepository) GetAll() ([]models.TrashCan, error) {
	trashCans := []models.TrashCan{}

	error := dbClient.Model(&trashCans).Where("1 = 1").Select()

	if error != nil {
		return trashCans, error
	}

	return trashCans, nil
}

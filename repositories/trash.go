package repositories

import (
	"github.com/Trash-Men/api-server/models"
)

type TrashRepository struct{}

func (_ TrashRepository) Create(photoUrl string, latitude, longitude float64, area string) error {
	_, error := dbClient.Model(&models.Trash{
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

func (_ TrashRepository) GetAll() ([]models.Trash, error) {
	trashes := []models.Trash{}

	error := dbClient.Model(&trashes).Where("1 = 1").Select()

	if error != nil {
		return trashes, error
	}

	return trashes, nil
}

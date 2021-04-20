package services

import (
	"github.com/Trash-Men/api-server/models"
)

type TrashService struct{}

func (_ TrashService) Create(photoUrl string, latitude, longitude float64, area string) error {
	return _repositories.TrashRepository.Create(photoUrl, latitude, longitude, area)
}

func (_ TrashService) GetAll() ([]models.Trash, error) {
	return _repositories.TrashRepository.GetAll()
}

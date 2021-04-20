package services

import (
	"github.com/Trash-Men/api-server/models"
)

type TrashCanService struct{}

func (_ TrashCanService) Create(photoUrl string, latitude, longitude float64, area string) error {
	return _repositories.TrashCanRepository.Create(photoUrl, latitude, longitude, area)
}

func (_ TrashCanService) GetAll() ([]models.TrashCan, error) {
	return _repositories.TrashCanRepository.GetAll()
}

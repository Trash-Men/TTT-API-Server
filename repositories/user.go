package repositories

import (
	"github.com/Trash-Men/api-server/models"
)

type UserRepository struct{}

func (_ UserRepository) CreateUser() error {
	_, error := dbClient.Model(&models.User{
		Id:       "cc",
		Password: "$2a$04$K4OzFE//n3jmGpGVKbRGe.7ibVuqo8Ak5TxRtf0OZkp4p6hxYLGGu",
		Role:     "admin",
	}).Insert()

	if error != nil {
		return error
	}

	return nil
}

func (_ UserRepository) GetUser(id, password string) (models.User, error) {
	emptyUser := models.User{}
	user := &models.User{
		Id: id,
	}

	error := dbClient.Model(user).WherePK().Select()

	if error != nil {
		return emptyUser, error
	}

	return *user, nil
}

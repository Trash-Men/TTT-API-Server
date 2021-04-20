package controllers

import (
	"github.com/Trash-Men/api-server/services"
	"github.com/Trash-Men/api-server/utils"
)

type ControllersStruct struct {
	UserController     UserController
	PhotoController    PhotoController
	TrashController    TrashController
	TrashCanController TrashCanController
}

func Controllers() ControllersStruct {
	return ControllersStruct{
		UserController:     UserController{},
		PhotoController:    PhotoController{},
		TrashController:    TrashController{},
		TrashCanController: TrashCanController{},
	}
}

var _services services.ServicesStruct = services.Services()

var errorCreater = utils.ErrorCreater{}

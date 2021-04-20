package controllers

import (
	"github.com/Trash-Men/api-server/models"
	"github.com/Trash-Men/api-server/services"
	"github.com/Trash-Men/api-server/utils"
	"github.com/labstack/echo"
)

type TrashCanController struct{}

func (_ TrashCanController) CreateTrashCan(context echo.Context) error {
	errorCreater.SetContext(context)

	createTrashRequest, error := utils.JsonMapper(context)

	if error != nil {
		return errorCreater.CreateCustomError(400, utils.ErrorCode("TC-01"), "고 코드에서 JSON 매핑하는 과정에서 실패.")
	}

	photoUrl, photoUrlOk := createTrashRequest["photoUrl"].(string)
	latitude, latitudeOk := createTrashRequest["latitude"].(float64)
	longitude, longitudeOk := createTrashRequest["longitude"].(float64)
	area, areaOk := createTrashRequest["area"].(string)

	if photoUrl == "" || latitude == 0.0 || longitude == 0.0 || area == "" || !photoUrlOk || !latitudeOk || !longitudeOk || !areaOk {
		return errorCreater.CreateCustomError(400, utils.ErrorCode("TC-02"), "all keys must be valid")
	}

	error = _services.TrashCanService.Create(photoUrl, latitude, longitude, area)

	if error != nil {
		return errorCreater.CreateCustomError(409, utils.ErrorCode("TC-03"), error.Error())
	}

	return context.NoContent(201)
}

func (_ TrashCanController) GetTrashCans(context echo.Context) error {
	trashCans, error := services.TrashCanService{}.GetAll()

	if error != nil {
		return errorCreater.CreateCustomError(409, utils.ErrorCode("TC-04"), "Failed Get Trash Cans")
	}

	return context.JSON(200, map[string][]models.TrashCan{
		"trashCans": trashCans,
	})
}

package controllers

import (
	"github.com/Trash-Men/api-server/models"
	"github.com/Trash-Men/api-server/services"
	"github.com/Trash-Men/api-server/utils"
	"github.com/labstack/echo"
)

type TrashController struct{}

func (_ TrashController) CreateTrash(context echo.Context) error {
	errorCreater.SetContext(context)

	createTrashRequest, error := utils.JsonMapper(context)

	if error != nil {
		return errorCreater.CreateCustomError(400, utils.ErrorCode("TH-01"), "고 코드에서 JSON 매핑하는 과정에서 실패.")
	}

	photoUrl, photoUrlOk := createTrashRequest["photoUrl"].(string)
	latitude, latitudeOk := createTrashRequest["latitude"].(float64)
	longitude, longitudeOk := createTrashRequest["longitude"].(float64)
	area, areaOk := createTrashRequest["area"].(string)

	if photoUrl == "" || latitude == 0.0 || longitude == 0.0 || area == "" || !photoUrlOk || !latitudeOk || !longitudeOk || !areaOk {
		return errorCreater.CreateCustomError(400, utils.ErrorCode("TH-02"), "all keys must be valid")
	}

	error = _services.TrashService.Create(photoUrl, latitude, longitude, area)

	if error != nil {
		return errorCreater.CreateCustomError(409, utils.ErrorCode("TH-03"), error.Error())
	}

	return context.NoContent(201)
}

func (_ TrashController) GetTrashes(context echo.Context) error {
	trashes, error := services.TrashService{}.GetAll()

	if error != nil {
		return errorCreater.CreateCustomError(409, utils.ErrorCode("TH-04"), "Failed Get Trashes")
	}

	return context.JSON(200, map[string][]models.Trash{
		"trashes": trashes,
	})
}

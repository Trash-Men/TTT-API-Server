package controllers

import (
	"github.com/Trash-Men/api-server/utils"
	"github.com/labstack/echo"
)

type PhotoController struct{}

func (photoController PhotoController) UploadPhoto(context echo.Context) error {
	errorCreater.SetContext(context)

	photoType := context.FormValue("type")
	photo, error := context.FormFile("photo")

	if photoType != "trash" && photoType != "trashCan" {
		return errorCreater.CreateCustomError(400, utils.ErrorCode("PH-01"), "photo type must be 'trash' or 'trashCan'")
	}

	if error != nil {
		return errorCreater.CreateCustomError(400, utils.ErrorCode("PH-02"), "photo key not provided")
	}

	photoPath, error := _services.PhotoService.UploadPhoto(photo, photoType)

	if error != nil {
		return errorCreater.CreateCustomError(400, utils.ErrorCode("PH-03"), error.Error())
	}

	return context.JSON(201, map[string]string{
		"photoPath": photoPath,
	})
}

func (photoController PhotoController) DeletePhoto(context echo.Context) error {
	errorCreater.SetContext(context)

	photoPath := context.QueryParam("photoPath")

	if photoPath == "" {
		return errorCreater.CreateCustomError(400, utils.ErrorCode("PH-04"), "'photoPath' must be required")
	}

	if error := _services.PhotoService.DeletePhoto(photoPath); error != nil {
		return errorCreater.CreateCustomError(400, utils.ErrorCode("PH-05"), error.Error())
	}

	return context.NoContent(200)
}

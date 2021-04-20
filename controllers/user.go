package controllers

import (
	"github.com/Trash-Men/api-server/utils"
	"github.com/labstack/echo"
)

type UserController struct{}

func (_ UserController) Login(context echo.Context) error {
	errorCreater.SetContext(context)

	loginRequest, error := utils.JsonMapper(context)

	if error != nil {
		return errorCreater.CreateCustomError(400, utils.ErrorCode("LO-01"), error.Error())
	}

	id, idOk := loginRequest["id"].(string)
	password, passwordOk := loginRequest["password"].(string)

	if !idOk || !passwordOk {
		return errorCreater.CreateCustomError(400, utils.ErrorCode("LO-02"), "body로부터 'id' 또는 'password' 값을 정상적으로 불러오지 못함.")
	}

	if id == "" || password == "" {
		return errorCreater.CreateCustomError(400, utils.ErrorCode("LO-03"), "'id'와 'password'는 빈 문자열 일 수 없음.")
	}

	token, error := _services.UserService.Login(id, password)

	type loginResponse struct {
		AccessToken string `json:"accessToken"`
	}

	if error != nil {
		return errorCreater.CreateCustomError(400, utils.ErrorCode("LO-04"), error.Error())
	}

	return context.JSON(200, map[string]string{
		"accessToken": token,
	})
}

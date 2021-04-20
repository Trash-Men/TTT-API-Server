package utils

import "github.com/labstack/echo"

// var errorCreaterInstance *errorCreater

type ErrorCreater struct {
	context echo.Context
}

// func ErrorCreaterInstance() *errorCreater {
// 	if errorCreaterInstance == nil {
// 		errorCreaterInstance = new(errorCreater)
// 	}
// 	return errorCreaterInstance
// }

func (creater *ErrorCreater) SetContext(context echo.Context) {
	creater.context = context
}

type ErrorCode string

type customError struct {
	Status  int       `json:"status"`
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func (creater ErrorCreater) CreateCustomError(status int, errorCode ErrorCode, errorMessage string) error {
	json := &customError{
		Status:  status,
		Code:    errorCode,
		Message: errorMessage,
	}

	return creater.context.JSON(status, json)
}

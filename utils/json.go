package utils

import (
	"encoding/json"

	"github.com/labstack/echo"
)

func JsonMapper(context echo.Context) (map[string]interface{}, error) {
	requestJson := make(map[string]interface{})
	error := json.NewDecoder(context.Request().Body).Decode(&requestJson)

	if error != nil {
		return nil, error
	}

	return requestJson, nil
}

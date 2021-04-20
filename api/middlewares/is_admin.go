package middlewares

import (
	"github.com/Trash-Men/api-server/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {

		token := context.Get("token").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		role := claims["role"].(string)

		if role == "admin" {
			return next(context)
		}

		errorCreater := utils.ErrorCreater{}
		errorCreater.SetContext(context)

		return errorCreater.CreateCustomError(403, "", "zzzz 어드민 아님?ㅋㅋ")
	}
}

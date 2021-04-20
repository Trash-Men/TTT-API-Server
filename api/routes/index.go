package routes

import (
	"github.com/Trash-Men/api-server/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var router *echo.Echo

func init() {
	router = echo.New()
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE},
	}))
}

func SetRouting() *echo.Echo {
	Photo()
	Trash()
	TrashCan()
	User()

	return router
}

var _controllers = controllers.Controllers()

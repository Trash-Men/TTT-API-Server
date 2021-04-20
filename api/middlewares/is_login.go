package middlewares

import (
	"github.com/Trash-Men/api-server/configs"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningMethod: "HS256",
	SigningKey:    []byte(configs.GetEnvironments().JWT_SECRET_KEY),
	ContextKey:    "token",
	TokenLookup:   "header:" + echo.HeaderAuthorization,
	AuthScheme:    "Bearer",
})

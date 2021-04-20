package routes

import (
	"github.com/Trash-Men/api-server/api/middlewares"
)

func Photo() {
	photoRouter := router.Group("/photo")

	photoRouter.Use(middlewares.IsLoggedIn)

	photoController := _controllers.PhotoController

	photoRouter.POST("", photoController.UploadPhoto)
	photoRouter.DELETE("", photoController.DeletePhoto, middlewares.IsAdmin)
}

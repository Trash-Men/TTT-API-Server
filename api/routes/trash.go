package routes

import "github.com/Trash-Men/api-server/api/middlewares"

func Trash() {
	trashRouter := router.Group("/trash")

	trashController := _controllers.TrashController

	trashRouter.Use(middlewares.IsLoggedIn)

	trashRouter.POST("", trashController.CreateTrash)
	trashRouter.GET(("/all"), trashController.GetTrashes)
}

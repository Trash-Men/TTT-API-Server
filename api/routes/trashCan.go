package routes

import "github.com/Trash-Men/api-server/api/middlewares"

func TrashCan() {
	trashCanRouter := router.Group("/trash-can")

	trashCanController := _controllers.TrashCanController

	trashCanRouter.Use(middlewares.IsLoggedIn)

	trashCanRouter.POST("", trashCanController.CreateTrashCan)
	trashCanRouter.GET("/all", trashCanController.GetTrashCans)
}

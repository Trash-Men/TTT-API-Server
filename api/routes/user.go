package routes

func User() {
	userRouter := router.Group("/user")

	UserController := _controllers.UserController

	userRouter.POST("/login", UserController.Login)
}

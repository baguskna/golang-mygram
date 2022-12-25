package router

import (
	"golang-mygram/app"
	"golang-mygram/controller"
	"golang-mygram/middleware"
	"golang-mygram/repository"
	"golang-mygram/service"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	app.StartDB()
	db := app.GetDB()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", userController.PostUser)
		userRouter.POST("/login", userController.LoginUser)
		
		userRouter.Use(middleware.Auth())
		userRouter.DELETE("/:id", userController.DeleteUser)
		userRouter.GET("/:id", userController.GetUserById)
	}

	return r
}

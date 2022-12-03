package router

import (
	"golang-mygram/app"
	"golang-mygram/controller"
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

	r.POST("/users/register", userController.PostUser)
	r.DELETE("/users/:id", userController.DeleteUser)
	r.GET("/users/:id", userController.GetUserById)

	return r
}

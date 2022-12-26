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

	photoRepository := repository.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository)
	photoController := controller.NewPhotoController(photoService)

	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository)
	commentController := controller.NewCommentController(commentService)

	socialMediaRepository := repository.NewSocialMediaRepository(db)
	socialMediaService := service.NewSocialMediaService(socialMediaRepository)
	socialMediaController := controller.NewSocialMediaController(socialMediaService)

	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", userController.PostUser)
		userRouter.POST("/login", userController.LoginUser)

		userRouter.Use(middleware.Auth())
		userRouter.GET("/", userController.GetUsers)
		userRouter.DELETE("/:id", userController.DeleteUser)
		userRouter.GET("/:id", userController.GetUserById)
		userRouter.PUT("/:id", userController.UpdateUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.GET("/", photoController.GetPhotos)
		photoRouter.Use(middleware.Auth())
		photoRouter.GET("/:id", photoController.GetPhotoById)
		photoRouter.POST("/", photoController.PostPhoto)
		photoRouter.DELETE("/:id", photoController.DeletePhoto)
		photoRouter.PUT("/:id", photoController.UpdatePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.GET("/", commentController.GetComments)
		photoRouter.Use(middleware.Auth())
		commentRouter.GET("/:id", commentController.GetCommentBtId)
		commentRouter.POST("/", commentController.PostComment)
		commentRouter.DELETE("/:id", commentController.DeleteComment)
		commentRouter.PUT("/:id", commentController.UpdateComment)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		// socialMediaRouter.Use(middleware.Authentication())
		socialMediaRouter.GET("/", socialMediaController.GetSocialMedias)
		socialMediaRouter.GET("/:id", socialMediaController.GetUserById)
		socialMediaRouter.POST("/", socialMediaController.PostSocialMedia)
		socialMediaRouter.DELETE("/:id", socialMediaController.DeleteSocialMedia)
		socialMediaRouter.PUT("/:id", socialMediaController.UpdateSocialMedia)
	}

	return r
}

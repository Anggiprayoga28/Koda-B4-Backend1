package routes

import (
	"backend1/controllers"
	"backend1/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(middleware.CORS())

	authController := controllers.NewAuthController()
	userController := controllers.NewUserController()
	profileController := controllers.NewProfileController()

	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	users := r.Group("/users")
	{
		users.GET("", userController.GetUsers)
		users.GET("/:id", userController.GetUserByID)
		users.POST("", userController.CreateUser)
		users.PATCH("/:id", userController.UpdateUser)
		users.DELETE("/:id", userController.DeleteUser)
	}

	upload := r.Group("/upload")
	{
		upload.POST("/:user_id", profileController.UploadPicture)
		upload.GET("/:user_id", profileController.GetPicture)
	}

	r.Static("/uploads", "./uploads")
}

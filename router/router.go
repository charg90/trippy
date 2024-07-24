package router

import (
	"trippy/middleware"
	controller "trippy/modules/user/user-controller"
	repositories "trippy/modules/user/user-repository/postgres-repository"
	services "trippy/modules/user/user.service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userRepository := repositories.NewPostgresUserRepository()
	userService := services.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	
	r.POST("/users", userController.CreateUser)
	r.POST("/login", userController.Login)
	r.GET("/users",middleware.AuthMiddleware(), userController.GetUsers)

	return r
}
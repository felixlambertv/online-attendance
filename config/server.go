package config

import (
	"fmt"
	"github.com/felixlambertv/online-attendance/controller"
	"github.com/felixlambertv/online-attendance/middleware"
	"github.com/felixlambertv/online-attendance/repository"
	"github.com/felixlambertv/online-attendance/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

func SetupServer(db *gorm.DB) *gin.Engine {
	appConfig := GetConfig()
	engine := gin.Default()
	engine.Use(gin.Recovery())
	engine.Use(middleware.ErrorHandler)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	v1Routes := engine.Group("api/v1")
	{
		v1Routes.GET("/ping", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"pong": appConfig.AppName})
		})

		userRoutes := v1Routes.Group("/users")
		{
			userRoutes.POST("/", userController.Register)
			userRoutes.GET("/:userId", userController.Detail)
		}
	}

	err := engine.Run(fmt.Sprintf(":%s", appConfig.Server.Port))
	if err != nil {
		logrus.Error("fail run app engine")
	}

	return engine
}

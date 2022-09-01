package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

func SetupServer(db *gorm.DB) *gin.Engine {
	appConfig := GetConfig()
	engine := gin.Default()

	apiGroup := engine.Group("api")
	{
		v1Routes := apiGroup.Group("/v1")
		{
			v1Routes.GET("/ping", func(context *gin.Context) {
				context.JSON(http.StatusOK, gin.H{"ping": appConfig.AppName})
			})
		}
	}
	err := engine.Run(fmt.Sprintf(":%s", appConfig.Server.Port))
	if err != nil {
		logrus.Error("fail run app engine")
	}

	return engine
}

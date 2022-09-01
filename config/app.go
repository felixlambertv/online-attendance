package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Server struct {
	Port string
}

type Database struct {
	Host     string
	Driver   string
	User     string
	Password string
	Name     string
	Port     string
}

type AppConfig struct {
	AppName  string
	Server   Server
	Database Database
}

func GetConfig() AppConfig {
	viper.SetConfigFile("yaml")
	viper.AddConfigPath("./")
	viper.SetConfigName("app.config")

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Error("failed read config")
	}

	server := Server{
		Port: viper.GetString("server.port"),
	}

	database := Database{
		Host:     viper.GetString("database.host"),
		Driver:   viper.GetString("database.driver"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Name:     viper.GetString("database.name"),
		Port:     viper.GetString("database.port"),
	}
	return AppConfig{
		AppName:  viper.GetString("appName"),
		Server:   server,
		Database: database,
	}
}

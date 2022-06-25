package main

import (
	"github.com/hunkevych-philip/mono-app/pkg/handler"
	"github.com/hunkevych-philip/mono-app/pkg/service"
	"github.com/hunkevych-philip/mono-app/pkg/utils"
	"github.com/hunkevych-philip/mono-app/server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("Could not read config file: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Cound not read env variables: %s", err.Error())
	}

	services := service.NewService()
	utilities := utils.NewUtils()
	handler := handler.NewHandler(services, utilities)

	s := new(server.Server)
	if err := s.Start(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("Server returned an error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

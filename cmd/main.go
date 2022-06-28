package main

import (
	"github.com/hunkevych-philip/mono-app/pkg/handler"
	"github.com/hunkevych-philip/mono-app/pkg/service"
	"github.com/hunkevych-philip/mono-app/pkg/utils"
	"github.com/hunkevych-philip/mono-app/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	ConfigKeyPort  string = "port"
	ConfigKeyDebug string = "debug"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("Could not read config file: %s", err.Error())
	}

	switch viper.GetBool(ConfigKeyDebug) {
	case true:
		logrus.SetLevel(logrus.DebugLevel)
	case false:
		logrus.SetLevel(logrus.InfoLevel)
	}

	services, err := service.NewService()
	if err != nil {
		logrus.Fatalf("Services initialization failed: %s", err.Error())
	}

	utilities := utils.NewUtils()
	handler := handler.NewHandler(services, utilities)

	s := new(server.Server)

	port := viper.GetString(ConfigKeyPort)
	if len(port) == 0 {
		logrus.Infof("%q is not set in a config file. Using default: 8080", ConfigKeyPort)
		port = "8080"
	}

	logrus.Infof("Starting server on a port: %s\n", port)
	if err := s.Start(port, handler.InitRoutes()); err != nil {
		logrus.Fatalf("Server returned an error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

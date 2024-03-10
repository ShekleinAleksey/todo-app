package main

import (
	"log"

	"github.com/ShekleinAleksey/todo-app.git"
	"github.com/ShekleinAleksey/todo-app.git/pkg/handler"
	"github.com/ShekleinAleksey/todo-app.git/pkg/repository"
	"github.com/ShekleinAleksey/todo-app.git/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

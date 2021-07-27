package main

import (
	"log"

	"github.com/semenov9480/hamser-api-golang/database"
	"github.com/semenov9480/hamser-api-golang/handler"
	"github.com/spf13/viper"
)

func main() {

	if err := InitConfig(); err != nil {
		log.Fatalf("Error init config: %s", err.Error())
	}

	_, err := database.MongoConnect(database.Config{
		Host: viper.GetString("DBhost"),
		Port: viper.GetString("DBport"),
	})
	if err != nil {
		log.Fatalf("Error connection db: %s", err.Error())
	}

	//запуск сервера
	handlers := new(handler.Handler)
	serv := new(handler.Server)
	if err := serv.Run(viper.GetString("ServerPort"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error starting https server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

package main

import (
	"log"

	"./handler"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("Error init config: %s", err.Error())
	}
	handlers := new(handler.Handler)
	serv := new(handler.Server)
	if err := serv.Run("8081", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error starting https server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

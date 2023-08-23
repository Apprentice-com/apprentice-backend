package main

import (
	"log"

	"github.com/KadirbekSharau/apprentice-backend/configs"
	"github.com/KadirbekSharau/apprentice-backend/internal/server"
	"github.com/spf13/viper"
)

func main() {
	if err := configs.Init(); err != nil {
		log.Fatal("%s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}

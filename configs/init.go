package configs

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func Init() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

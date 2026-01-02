package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	if err := godotenv.Load(".env"); err != nil {
		panic(".env file not found")
		os.Exit(1)
	}
	config := viper.New()
	config.AutomaticEnv()

	return config
}

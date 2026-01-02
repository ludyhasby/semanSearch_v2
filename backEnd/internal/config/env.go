package config

import "github.com/spf13/viper"

type Env struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
}

func NewEnv(viper *viper.Viper) *Env {
	return &Env{
		DBUser: viper.GetString("DB_USER"),
		DBPass: viper.GetString("DB_PASS"),
		DBHost: viper.GetString("DB_HOST"),
		DBPort: viper.GetString("DB_PORT"),
		DBName: viper.GetString("DB_NAME"),
	}
}

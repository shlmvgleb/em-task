package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Port int
}

func ReadEnvConfig() AppConfig {
	viper.SetConfigFile("./.env")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return AppConfig{}
}

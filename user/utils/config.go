package utils

import (
	"log"

	"github.com/spf13/viper"
)

type serviceConfig struct {
	Port        string `mapstructure:"port"`
	DatabaseURL string `mapstructure:"database_url"`
	OrderURL    string `mapstructure:"order_host"`
}

var sc serviceConfig

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln(err)
	}
	err = viper.Unmarshal(&sc)
	if err != nil {
		log.Panicln(err)
	}
}

func GetServiceConfig() *serviceConfig {
	return &sc
}

package cmd

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file : %w", err))
	}

}

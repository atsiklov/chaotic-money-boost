package config

import (
	"log"

	"github.com/spf13/viper"
)

func loadYaml(appCfg *AppConfig) {
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error while reading config: " + err.Error())
	}
	if err := viper.Unmarshal(appCfg); err != nil {
		log.Fatal("Error while parsing config: " + err.Error())
	}
}

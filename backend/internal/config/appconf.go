package config

import "fmt"

type HTTPServer struct {
	Port string `mapstructure:"port"`
}

type DbServer struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
	Name string `mapstructure:"name"`
}

func (config DbServer) DbConnectionParams() string {
	return fmt.Sprintf(
		"user=%s host=%s port=%s dbname=%s",
		config.User, config.Host, config.Port, config.Name,
	)
}

type AppConfig struct {
	HTTPServer HTTPServer `mapstructure:"http_server"`
	DbServer   DbServer   `mapstructure:"db_server"`
}

var AppConf AppConfig

func Load() {
	// 1 загружаем переменные из файла .env в оперативу (в формате key: value)
	loadEnvToMemory()

	// 2 маппим параметры и значения из config.yaml в объект AppConfig
	loadYaml(&AppConf)

	// 3 перезатираем поля в AppConfig параметрами из .env (которые загрузили в п.1)
	overrideWithEnv(&AppConf)
}

package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port     string `json:"port"`
	Database string `json:"database"`
}

func Init() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file %v", err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode config %v", err)
	}

	return config
}

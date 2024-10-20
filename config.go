package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Servers []Server `mapstructure:"servers"`
}

func LoadConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, fmt.Errorf("error reading config: %v", err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config: %v", err)
	}

	return config, nil
}

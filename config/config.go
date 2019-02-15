package config

import (
	"github.com/davyj0nes/stubby"
	"github.com/spf13/viper"
)

type Config struct {
	Routes []stubby.Route
}

func LoadConfig(fileName string) (Config, error) {
	var cfg Config

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(fileName)

	err := v.ReadInConfig()
	if err != nil {
		return cfg, err
	}

	err = v.Unmarshal(&cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, err
}

package config

import (
	"github.com/davyj0nes/stubby/internal/router"
	"github.com/spf13/viper"
)

// Config is used to marshal a configuration file into something more useful
type Config struct {
	Port   int
	Routes []router.Route
}

// LoadConfig from a file
func LoadConfig(fileName string) (Config, error) {
	var cfg Config

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(fileName)

	err := v.ReadInConfig()
	if err != nil {
		return cfg, err
	}

	if err = v.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

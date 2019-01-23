package config

import "github.com/spf13/viper"

type Route struct {
	Path     string
	Response string
}

type Config struct {
	Port   string
	Routes []Route
}

func LoadConfig(fileName string) (Config, error) {
	var cfg Config
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(fileName)
	v.AddConfigPath("./")
	v.AutomaticEnv()

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

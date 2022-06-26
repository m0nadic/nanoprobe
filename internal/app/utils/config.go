package utils

import (
	"github.com/spf13/viper"
	"nanoprobe/internal/app/configs"
)

func LoadConfig(cfgFile string) (*configs.Config, error) {
	viper.SetConfigFile(cfgFile)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var cfg configs.Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, err
}

package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

var globalConfig *Config

func init() {
	globalConfig = &Config{
		viper.New(),
	}
	globalConfig.SetConfigType("yml")
	//globalConfig.SetConfigName("config")
	globalConfig.SetConfigName("config")
	globalConfig.AddConfigPath(".")
	if err := globalConfig.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	logrus.SetLevel(logrus.DebugLevel)
}

func GetConfig() *Config {
	return globalConfig
}

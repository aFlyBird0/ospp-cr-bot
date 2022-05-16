package config

import (
	"fmt"

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
}

func GetConfig() *Config {
	return globalConfig
}

func IsProd() bool {
	return globalConfig.GetString("env") == "prod"
}

func IsGitPlatformEnabled(platform string) bool {
	return globalConfig.GetBool(fmt.Sprintf("platforms.git.%s.enabled", platform))
}

func IsCommunityEnabled(community string) bool {
	return globalConfig.GetBool(fmt.Sprintf("platforms.community.%s.enabled", community))
}

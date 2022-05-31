package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func IsProd() bool {
	return globalConfig.GetString("env") == "prod"
}

func GetGitPlatformConfig(platform string) *viper.Viper {
	return globalConfig.Sub("platforms.git." + platform)
}

func GetCommunityPlatformConfig(platform string) *viper.Viper {
	return globalConfig.Sub("platforms.community." + platform)
}

func IsGitPlatformEnabled(platform string) bool {
	return globalConfig.GetBool(fmt.Sprintf("platforms.git.%s.enabled", platform))
}

func IsCommunityEnabled(community string) bool {
	return globalConfig.GetBool(fmt.Sprintf("platforms.community.%s.enabled", community))
}

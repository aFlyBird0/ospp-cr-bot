package union

import "github.com/devstream-io/devstream/ospp-cr-bot/config"

type (
	// UserConfig is user defined in config
	UserConfig struct {
		Git       map[string]string
		Community map[string]string
	}

	UserConfigs struct {
		Users []*UserConfig
	}
)

func getUserConfigs() []*UserConfig {
	userConfigs := UserConfigs{}
	if err := config.GetConfig().Unmarshal(&userConfigs); err != nil {
		panic(err)
	}
	return userConfigs.Users
}

type (
	// RepoConfig is repo defined in config
	RepoConfig struct {
		Platform string
		Name     string
		Groups   map[string]string
		Admins   []string
		Users    []string
	}

	RepoConfigs struct {
		Repos []*RepoConfig
	}
)

func getRepoConfigs() []*RepoConfig {
	repoConfigs := RepoConfigs{}
	if err := config.GetConfig().Unmarshal(&repoConfigs); err != nil {
		panic(err)
	}
	return repoConfigs.Repos
}

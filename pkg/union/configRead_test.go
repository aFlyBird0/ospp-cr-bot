package union

import (
	"fmt"
	"github.com/devstream-io/devstream/ospp-cr-bot/config"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/community"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"
	"testing"
)

func TestReadUserConfig(t *testing.T) {
	users := getUserConfigs()
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}
}

func TestReadRepoConfig(t *testing.T) {
	repos := getRepoConfigs()
	for _, repo := range repos {
		fmt.Printf("%+v\n", repo)
	}
}

func TestPlatformsEnabled(t *testing.T) {
	if config.IsGitPlatformEnabled(string(git.PlatformGithub)) {
		t.Log("Github is enabled")
	} else {
		t.Error("Github is not enabled")
	}

	if config.IsCommunityEnabled(string(community.CommunityTypeFeishu)) {
		t.Log("Feishu is enabled")
	} else {
		t.Error("Feishu is not enabled")
	}
}

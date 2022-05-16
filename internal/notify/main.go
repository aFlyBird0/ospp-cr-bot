package notify

import (
	"fmt"
	"github.com/devstream-io/devstream/ospp-cr-bot/config"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/community"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/union"
	"time"

	_ "github.com/devstream-io/devstream/ospp-cr-bot/internal/community/feishuMock" // important, to call init() to register client
	_ "github.com/devstream-io/devstream/ospp-cr-bot/internal/git/githubMock"       // important, to call init() to register client
)

func Main() {
	union.Init()

	gitPlatforms := git.GetPlatformMap()
	for _, platform := range gitPlatforms {
		fmt.Printf("git platform [%s] is registered\n", platform.GetType())
	}

	communities := community.GetCommunityMap()
	for _, c := range communities {
		fmt.Printf("community [%s] is registered\n", c.GetType())
	}

	// execute once in debug mode
	if !config.IsProd() {
		RefreshCommunities()
		NotifyPr()
	}

	go Polling()

}

// Polling 轮询
func Polling() {
	ticker := time.NewTicker(time.Minute * 5)
	for {
		select {
		case <-ticker.C:
			fmt.Println("Tick")
			RefreshCommunities()
			NotifyPr()
		}
	}
}

func NotifyPr() {
	//issueOpenFilter := git.IssueFilter{State: git.IssueStateOpen}
	prOpenFilter := git.PrFilter{State: git.PrStateOpen}

	// todo make notify duration configurable
	// todo record last notified time
	// todo record admins' response
	// todo auto change duration based on admins' response
	// todo support more activities

	// todo need refactor this process
	for _, platform := range git.GetPlatformMap() {
		for _, repo := range union.ListRepoByGitPlatform(platform.GetType()) {
			prs, err := platform.ListPrsByRepoWithFilter(repo.GitRepo, prOpenFilter)
			if err != nil {
				fmt.Println(err)
				continue
			}
			for _, pr := range prs {
				messages := union.AnalysePrCausedByWho(pr)
				for _, message := range messages {
					for _, c := range community.GetCommunityMap() {
						communityUser, ok := union.GetCommunityUserByGitUser(message.GetToGitUser(), c.GetType())
						if !ok {
							fmt.Printf("notify error: git user [%s] not found of PR(#%d) in %v", message.GetToGitUser().GetUserID(), pr.GetNumber(), c.GetType())
							continue
						}
						err := c.SendMessageToUsersInGroup(&message, []community.User{communityUser}, union.GetGroupByRepoAndCommunity(repo.GitRepo, c.GetType()))
						if err != nil {
							fmt.Println(err)
						}
					}
				}
			}

		}
	}
}

func RefreshCommunities() {
	for _, c := range community.GetCommunityMap() {
		if refresh, ok := c.(community.NeedRefresh); ok {
			if err := refresh.Refresh(); err != nil {
				fmt.Println(err)
			}
		}
	}
}

package notify

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/devstream-io/devstream/pkg/util/log"
	"github.com/sirupsen/logrus"

	"github.com/devstream-io/devstream/ospp-cr-bot/config"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/community"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/union"

	_ "github.com/devstream-io/devstream/ospp-cr-bot/internal/pkg/community/feishuMock" // important, to call init() to register client
	_ "github.com/devstream-io/devstream/ospp-cr-bot/internal/pkg/git/github"           // important, to call init() to register client
)

func init() {
	if !config.IsProd() {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func Main() {
	union.Init()

	// execute once in debug mode
	if !config.IsProd() {
		RefreshCommunities()
		NotifyPr()
	}

	go Polling()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
}

// Polling 轮询
func Polling() {
	ticker := time.NewTicker(time.Minute * 5)
	for {
		select {
		case <-ticker.C:
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
	// todo record users' response
	// todo auto change duration based on users' response
	// todo support more activities

	// todo refactor process check with stream

	// todo need refactor this process
	for _, platform := range git.GetPlatformMap() {
		for _, repo := range union.ListRepoByGitPlatform(platform.GetType()) {
			log.Debugf("start polling repo: %v.", repo.GitRepo.GetName())
			prs, err := platform.ListPrsByRepo(repo.GitRepo, prOpenFilter)
			if err != nil {
				log.Errorf("list prs failed of repo [%v] from [%v]: %v.", repo.GitRepo.GetName(), repo.GitRepo.OfPlatForm().GetType(), err)
				continue
			}
			for _, pr := range prs {
				log.Debugf("get open pr: [%v].", pr.GetTitle())
				messages := union.AnalysePrCausedByWho(pr)
				for _, message := range messages {
					for _, c := range community.GetCommunityMap() {
						communityUser, ok := union.GetCommunityUserByGitUser(message.GetToGitUser(), c.GetType())
						if !ok {
							log.Errorf("notify error: git user [%s] not found in PR(#%d) from %v.", message.GetToGitUser().GetUserID(), pr.GetNumber(), c.GetType())
							continue
						}
						err := c.SendMessageToUserInGroup(&message, communityUser, union.GetGroupByRepoAndCommunity(repo.GitRepo, c.GetType()))
						if err != nil {
							log.Errorf("notify error: %v", err)
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
				log.Errorf("refresh community [%s] error: %v", c.GetType(), err)
			}
		}
	}
}

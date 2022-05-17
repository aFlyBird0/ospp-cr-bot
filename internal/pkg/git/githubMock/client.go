package githubMock

import (
	"github.com/devstream-io/devstream/ospp-cr-bot/config"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"
)

var client *Client

type Client struct {
}

func init() {
	if config.IsGitPlatformEnabled(string(git.PlatformGithub)) {
		client = &Client{}
		git.RegisterPlatform(client)
	}
}

func (c *Client) GetType() git.PlatformType {
	return git.PlatformGithub
}

func (c *Client) GetUserInfoByID(id string) (git.User, error) {
	return &User{Login: id}, nil
}

func (c *Client) GetRepoInfo(repoName string) (git.Repo, error) {
	// just mock, this func should call GitHub api to get more repo info
	return Repo{Name: repoName}, nil
}

func (c *Client) ListRepos() []git.Repo {
	return []git.Repo{
		Repo{Name: "repo1"},
		Repo{Name: "repo2"},
	}
}

func (c *Client) ListIssuesByRepo(repo git.Repo, filter git.IssueFilter) ([]git.Issue, error) {
	return []git.Issue{
		&Issue{Title: "issue1", Body: "body1"},
		&Issue{Title: "issue2", Body: "body2"},
	}, nil
}

func (c *Client) ListPrsByRepo(repo git.Repo, filter git.PrFilter) ([]git.PullRequest, error) {
	return []git.PullRequest{
		&PR{Title: "pr1", Body: "body1", Assignees: []*User{&User{Login: "github-user-1"}}},
		&PR{Title: "pr2", Body: "body2", Assignees: []*User{&User{Login: "github-user-2"}}},
	}, nil
}

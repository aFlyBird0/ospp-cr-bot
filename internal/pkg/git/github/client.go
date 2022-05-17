package github

import (
	ghUtil "github.com/devstream-io/devstream/pkg/util/github"
	"github.com/devstream-io/devstream/pkg/util/log"
	github "github.com/google/go-github/v42/github"

	"github.com/devstream-io/devstream/ospp-cr-bot/config"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"
)

// the way to create client is copy from devstream-io/devstream

var client *Client

type Client struct {
	*ghUtil.Client
}

func init() {
	if config.IsGitPlatformEnabled(string(git.PlatformGithub)) {
		option := buildOptionFromConfig()
		c, err := ghUtil.NewClient(option)
		if err != nil {
			log.Fatalf("failed to create github client: %v", err)
		}
		client = &Client{Client: c}
		git.RegisterPlatform(client)
	}
}

func buildOptionFromConfig() *ghUtil.Option {
	owner := config.GetConfig().GetString("platforms.git.github.owner")
	if owner == "" {
		log.Fatal("github owner is not set")
	}
	option := &ghUtil.Option{
		Owner: owner,
		// todo public repo do not need auth
		NeedAuth: true,
	}
	return option
}

func (c *Client) GetType() git.PlatformType {
	return git.PlatformGithub
}

func (c *Client) GetRepoInfo(repoName string) (git.Repo, error) {
	// note: to show how interface works
	repo, _, err := c.Repositories.Get(c.Context, c.Owner, repoName)
	if err != nil {
		return nil, err
	}
	return toRepoInf(repo), nil
}

func (c *Client) ListRepos() []git.Repo {
	// todo return errs
	repos, _, err := c.Repositories.List(c.Context, c.Owner, nil)
	if err != nil {
		log.Errorf("failed to list repos: %v", err)
		return nil
	}
	var ret []git.Repo
	for _, repo := range repos {
		ret = append(ret, toRepoInf(repo))
	}
	return ret
}

func (c *Client) ListIssuesByRepo(repo git.Repo, filter git.IssueFilter) ([]git.Issue, error) {
	issues, _, err := c.Issues.ListByRepo(c.Context, c.Owner, repo.GetName(), &github.IssueListByRepoOptions{
		State: issueStateToString(filter.State),
	})
	if err != nil {
		return nil, err
	}
	var ret []git.Issue
	for _, issue := range issues {
		ret = append(ret, toIssueInf(issue))
	}
	return ret, nil
}

func (c *Client) ListPrsByRepo(repo git.Repo, filter git.PrFilter) ([]git.PullRequest, error) {
	options := &github.PullRequestListOptions{
		State: prStateToString(filter.State),
	}
	prs, _, err := c.PullRequests.List(c.Context, c.Owner, repo.GetName(), options)
	if err != nil {
		return nil, err
	}
	var ret []git.PullRequest
	for _, pr := range prs {
		ret = append(ret, toPrInf(pr))
	}
	return ret, nil
}

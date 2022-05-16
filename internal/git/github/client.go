package github

import (
	"context"
	"fmt"
	"github.com/devstream-io/devstream/ospp-cr-bot/config"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"
	"github.com/devstream-io/devstream/pkg/util/log"
	"github.com/google/go-github/v44/github"
	"golang.org/x/oauth2"
	"os"
)

// this way to found client is copy from devstream-io/devstream

var client *Client

type Client struct {
	*github.Client
	*Option
	context.Context
}

func init() {
	if config.IsGitPlatformEnabled(string(git.PlatformGithub)) {
		// todo implement github client
		var err error
		if client, err = newClient(nil); err != nil {
			log.Fatalf("failed to create github client: %v", err)
		}
		git.RegisterPlatform(client)
	}
}

type Option struct {
	Owner    string
	Org      string
	Repo     string
	NeedAuth bool
}

func newClient(option *Option) (*Client, error) {
	if option == nil {
		option = &Option{}
	}
	// a. client without auth enabled
	if !option.NeedAuth {
		log.Debug("Auth is not enabled.")
		client = &Client{
			Option:  option,
			Client:  github.NewClient(nil),
			Context: context.Background(),
		}

		return client, nil
	}
	log.Debug("Auth is enabled.")

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		// github_token works well as GITHUB_TOKEN.
		token = os.Getenv("github_token")
	}
	if token == "" {
		retErr := fmt.Errorf("environment variable GITHUB_TOKEN is not set. Failed to initialize GitHub token. More info - " +
			"https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token")
		return nil, retErr
	}
	log.Debugf("Token: %s.", token)

	ctx := context.Background()
	tc := oauth2.NewClient(
		context.TODO(),
		oauth2.StaticTokenSource(
			&oauth2.Token{
				AccessToken: token,
			},
		),
	)

	client = &Client{
		Option:  option,
		Client:  github.NewClient(tc),
		Context: ctx,
	}

	return client, nil
}

func (c *Client) GetType() git.PlatformType {
	return git.PlatformGithub
}

func (c *Client) GetRepoInfo(repoName string) (git.Repo, error) {
	// TODO implement
	// note: to show how interface works
	repo := &github.Repository{}
	return toRepoInf(repo), nil
}

func (c *Client) ListRepos() []git.Repo {
	//TODO implement me
	return nil
}

func (c *Client) ListIssuesByRepo(repo git.Repo) ([]git.Issue, error) {
	//TODO implement me
	return nil, nil
}

func (c *Client) ListIssuesByRepoWithFilter(repo git.Repo, filter git.IssueFilter) ([]git.Issue, error) {
	//TODO implement me
	return nil, nil
}

func (c *Client) ListPrsByRepo(repo git.Repo) ([]git.PullRequest, error) {
	//TODO implement me
	return nil, nil
}

func (c *Client) ListPrsByRepoWithFilter(repo git.Repo, filter git.PrFilter) ([]git.PullRequest, error) {
	//TODO implement me
	return nil, nil
}

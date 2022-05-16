package github

import (
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"

	"github.com/google/go-github/v42/github"
)

type Repo struct {
	Name        string
	URL         string
	Description string
}

func toRepoInf(repository *github.Repository) Repo {
	return Repo{
		Name:        *repository.Name,
		URL:         *repository.HTMLURL,
		Description: *repository.Description,
	}
}

func (r Repo) GetName() string {
	return r.Name
}

func (r Repo) GetDescription() string {
	return r.Description
}

func (r Repo) GetURL() string {
	return r.URL
}

func (r Repo) OfPlatForm() git.Platform {
	return git.GetPlatformByType(git.PlatformGithub)
}

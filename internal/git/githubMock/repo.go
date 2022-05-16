package githubMock

import "github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"

type Repo struct {
	Name        string
	URL         string
	Description string
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

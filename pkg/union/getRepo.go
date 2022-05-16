package union

import (
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/community"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"
)

func ListAllRepos() []*Repo {
	return allRepos
}

func ListRepoByGitPlatform(platformType git.PlatformType) []*Repo {
	var repos []*Repo
	for _, repo := range allRepos {
		if repo.GitRepo.OfPlatForm().GetType() == platformType {
			repos = append(repos, repo)
		}
	}
	return repos
}

func GetRepoByGitRepo(repo git.Repo) (*Repo, bool) {
	for _, v := range allRepos {
		if SameRepo(v.GitRepo, repo) {
			return v, true
		}
	}
	return nil, false
}

func GetRepoByGroup(group community.Group) (*Repo, bool) {
	for _, v := range allRepos {
		for _, g := range v.Groups {
			if g.GetGroupID() == group.GetGroupID() {
				return v, true
			}
		}
	}
	return nil, false
}

func SameRepo(repo1, repo2 git.Repo) bool {
	return repo1.OfPlatForm() == repo2.OfPlatForm() && repo1.GetName() == repo2.GetName()
}

package union

import (
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/community"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"
)

func GetGitRepoGroups(repo git.Repo) []community.Group {
	var groups []community.Group
	for _, v := range allRepos {
		if SameRepo(v.GitRepo, repo) {
			groups = append(groups, v.Groups...)
		}
	}
	return groups
}

func GetGroupByRepoAndCommunity(repo git.Repo, communityType community.Type) community.Group {
	repoUnion, ok := GetRepoByGitRepo(repo)
	if !ok {
		return nil
	}
	for _, g := range repoUnion.Groups {
		if g.OfCommunity().GetType() == communityType {
			return g
		}
	}
	return nil
}

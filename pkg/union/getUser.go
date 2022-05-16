package union

import (
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/community"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"
)

func GetUnionUserByGitUser(user git.User) (*User, bool) {
	for _, u := range allUsers {
		for _, gitUser := range u.GitUsers {
			if gitUser.GetUserID() == user.GetUserID() {
				return u, true
			}
		}
	}
	return nil, false
}

func GetUnionUserByCommunityUser(user community.User) (*User, bool) {
	for _, u := range allUsers {
		for _, communityUser := range u.CommunityUsers {
			if communityUser.GetUserID() == user.GetUserID() {
				return u, true
			}
		}
	}
	return nil, false
}

func GetCommunityUserByGitUser(user git.User, communityType community.Type) (community.User, bool) {
	userUnion, ok := GetUnionUserByGitUser(user)
	if !ok {
		return nil, false
	}
	for _, communityUser := range userUnion.CommunityUsers {
		if communityUser.OfCommunity().GetType() == communityType {
			return communityUser, true
		}
	}
	return nil, false
}

func ListRepoGitAdmins(repo git.Repo) []git.User {
	for _, v := range allRepos {
		if SameRepo(v.GitRepo, repo) {
			return v.Admins
		}
	}
	return nil
}

func ListRepoGitUsers(repo git.Repo) []git.User {
	for _, v := range allRepos {
		if SameRepo(v.GitRepo, repo) {
			return v.Users
		}
	}
	return nil
}

func ListRepoUnionAdmins(repo git.Repo) UserList {
	for _, v := range allRepos {
		if SameRepo(v.GitRepo, repo) {
			users := make(UserList, 0)
			for _, u := range v.Admins {
				if user, ok := GetUnionUserByGitUser(u); ok {
					users = append(users, user)
				}
			}
			return users
		}
	}
	return nil
}

func ListRepoUnionUsers(repo git.Repo) UserList {
	for _, v := range allRepos {
		if SameRepo(v.GitRepo, repo) {
			users := make(UserList, 0)
			for _, u := range v.Users {
				if user, ok := GetUnionUserByGitUser(u); ok {
					users = append(users, user)
				}
			}
			return users
		}
	}
	return nil
}

func (l UserList) CommunityUsers() []community.User {
	var users []community.User
	for _, v := range l {
		users = append(users, v.CommunityUsers...)
	}
	return users
}

func (l UserList) GitUsers() []git.User {
	var users []git.User
	for _, v := range l {
		users = append(users, v.GitUsers...)
	}
	return users
}

package union

import (
	"fmt"

	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/community"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"

	"github.com/devstream-io/devstream/pkg/util/log"
)

// union interfaces from other packages

var (
	allUsers UserList
	allRepos RepoList
)

type (
	// User combines users from different sources
	User struct {
		GitUsers       []git.User
		CommunityUsers []community.User
	}
	UserList []*User

	// Repo combines repos from different sources with groups and users
	Repo struct {
		GitRepo git.Repo
		// groups that this repo belongs to, one repo only belongs to one group of each community type
		Groups []community.Group
		Admins []git.User // admins of the repo
		Users  []git.User // users who have access to this repo including admins
	}
	RepoList []*Repo
)

func Init() {

	errs := registerUsersFromConfig()
	for _, err := range errs {
		log.Warnf("Error during init users: %v", err)
	}

	errs, panics := initReposAndGroups()
	for _, err := range errs {
		log.Errorf("Error during init repos and groups : %v", err)
	}

	for _, panicErr := range panics {
		log.Fatalf("Fatal Error during init repos and groups : %v", panicErr)
	}

}

func registerUsersFromConfig() (errors []error) {
	usersFromConfig := getUserConfigs()
	for _, user := range usersFromConfig {
		gitUsers := make([]git.User, 0)
		communityUsers := make([]community.User, 0)
		for p, id := range user.Git {
			if u, err := git.RegisterUser(git.PlatformType(p), id); err != nil {
				errors = append(errors, err)
			} else {
				gitUsers = append(gitUsers, u)
			}
		}
		for c, id := range user.Community {
			if u, err := community.RegisterUser(community.Type(c), id); err != nil {
				errors = append(errors, err)
			} else {
				communityUsers = append(communityUsers, u)
			}
		}
		if len(gitUsers) > 0 || len(communityUsers) > 0 {
			allUsers = append(allUsers, &User{
				GitUsers:       gitUsers,
				CommunityUsers: communityUsers,
			})
		}
	}
	return
}

func initReposAndGroups() (errs []error, panicErrs []error) {
	reposFromConfig := getRepoConfigs()
	for _, repo := range reposFromConfig {

		if preCheckErrs := preCheckRepoConfig(repo); len(preCheckErrs) > 0 {
			errs = append(errs, preCheckErrs...)
			continue
		}

		groups := make([]community.Group, 0)
		admins := make([]git.User, 0)
		users := make([]git.User, 0)

		for g, id := range repo.Groups {
			if g, err := community.RegisterGroup(community.Type(g), id); err != nil {
				errs = append(errs, err)
			} else {
				groups = append(groups, g)
			}
		}
		for _, id := range repo.Admins {
			if u, ok := git.GetUserByTypeAndID(git.PlatformType(repo.Platform), id); !ok {
				errs = append(errs, fmt.Errorf("admin %s in repo %s not registered", id, repo.Name))
			} else {
				admins = append(admins, u)
			}
		}
		for _, id := range repo.Users {
			if u, ok := git.GetUserByTypeAndID(git.PlatformType(repo.Platform), id); !ok {
				errs = append(errs, fmt.Errorf("admin %s in repo %s not registered", id, repo.Name))
			} else {
				admins = append(admins, u)
			}
		}
		users = append(users, admins...)

		if len(groups) > 0 && len(users) > 0 {
			gitRepo, err := git.GetPlatformByType(git.PlatformType(repo.Platform)).GetRepoInfo(repo.Name)
			if err != nil {
				panicErrs = append(panicErrs, fmt.Errorf("get repo %s of %s info error: %v", repo.Name, repo.Platform, err))
				continue
			}
			allRepos = append(allRepos, &Repo{
				GitRepo: gitRepo,
				Groups:  groups,
				Admins:  admins,
				Users:   users,
			})
		} else {
			errs = append(errs, fmt.Errorf("repo %s has no groups or users", repo.Name))
		}
	}
	return
}

func preCheckRepoConfig(repo *RepoConfig) (errs []error) {
	// pre-check
	if repo.Platform == "" || repo.Name == "" {
		errs = append(errs, fmt.Errorf("repo config error: platform or name is empty"))
	}
	if len(repo.Groups) == 0 {
		errs = append(errs, fmt.Errorf("repo %s has no groups", repo.Name))
	}
	if len(repo.Users) == 0 {
		errs = append(errs, fmt.Errorf("repo %s has no users", repo.Name))
	}
	return errs
}

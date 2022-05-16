package github

import (
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"

	"github.com/google/go-github/v42/github"
)

type User struct {
	Login string
}

func toUserInf(user *github.User) *User {
	return &User{
		Login: user.GetLogin(),
	}
}

func (u *User) GetUserID() string {
	return u.Login
}

func (u *User) OfPlatform() git.Platform {
	return client
}

package github

import (
	"github.com/google/go-github/v42/github"

	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"
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

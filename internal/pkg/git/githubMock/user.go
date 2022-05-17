package githubMock

import (
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"
)

type User struct {
	Login string
}

func (u *User) GetUserID() string {
	return u.Login
}

func (u *User) OfPlatform() git.Platform {
	return client
}

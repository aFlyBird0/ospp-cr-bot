package feishuMock

import (
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/community"
)

type User struct {
	ID        string
	Name      string
	AvatarURL string
}

func (u *User) GetUserID() string {
	return u.ID
}

func (u *User) OfCommunity() community.Community {
	return client
}

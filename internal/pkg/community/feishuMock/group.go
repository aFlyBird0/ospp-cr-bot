package feishuMock

import "github.com/devstream-io/devstream/ospp-cr-bot/pkg/community"

type Group struct {
	ID   string
	Name string
}

func (g *Group) GetGroupID() string {
	return g.ID
}

func (g *Group) OfCommunity() community.Community {
	return client
}

package feishu

import "github.com/devstream-io/devstream/ospp-cr-bot/pkg/community"

type Group struct {
}

func (g *Group) GetGroupID() string {
	//TODO implement me
	panic("implement me")
}

func (g *Group) OfCommunity() community.Community {
	return client
}

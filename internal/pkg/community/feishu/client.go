package feishu

import (
	"github.com/devstream-io/devstream/ospp-cr-bot/config"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/community"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/message"
)

var client *Client

const CommunityTypeFeishu community.Type = "feishu"

type Client struct {
}

func init() {
	if config.IsCommunityEnabled(string(CommunityTypeFeishu)) {
		client = &Client{}
		community.RegisterCommunity(client)
	}
}

func (f Client) GetType() community.Type {
	return CommunityTypeFeishu
}

func (f Client) GetUserByID(userID string) (community.User, error) {
	//TODO implement me
	panic("implement me")
}

func (f Client) GetGroupByID(groupID string) (community.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (f Client) MessageFromIssue(issue git.Issue) message.Message {
	//TODO implement me
	panic("implement me")
}

func (f Client) MessageFromPr(pr git.PullRequest) message.Message {
	//TODO implement me
	panic("implement me")
}

func (f Client) MessageFromComment(comment git.Comment) message.Message {
	//TODO implement me
	panic("implement me")
}

func (f Client) SendMessageToUsers(message message.MessageInf, users []community.User) error {
	//TODO implement me
	panic("implement me")
}

func (f Client) SendMessageToGroup(message message.MessageInf, group community.Group) error {
	//TODO implement me
	panic("implement me")
}

func (f Client) SendMessageToUsersInGroup(message message.MessageInf, users []community.User, group community.Group) error {
	//TODO implement me
	panic("implement me")
}
func (f Client) SendMessageToUserInGroup(message message.MessageInf, users community.User, group community.Group) error {
	//TODO implement me
	panic("implement me")
}

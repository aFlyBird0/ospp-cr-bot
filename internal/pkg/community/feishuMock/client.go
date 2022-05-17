package feishuMock

import (
	"github.com/devstream-io/devstream/pkg/util/log"

	"github.com/devstream-io/devstream/ospp-cr-bot/config"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/community"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/message"
)

var client *Client

type Client struct {
}

func init() {
	if config.IsCommunityEnabled(string(community.CommunityTypeFeishu)) {
		client = &Client{}
		community.RegisterCommunity(client)
	}
}

func (c *Client) GetType() community.Type {
	return community.CommunityTypeFeishu
}

func (c *Client) GetUserByID(userID string) (community.User, error) {
	return nil, nil
}

func (c *Client) GetGroupByID(groupID string) (community.Group, error) {
	return nil, nil
}

func (c *Client) SendMessageToUsers(message message.MessageInf, users []community.User) error {
	// todo use go template
	log.Infof("SendMessageToUsers: %v\n", users)
	log.Infof("message: title: %v, body: %v, url: %v\n", message.GetTitle(), message.GetBody(), message.GetURL())
	return nil
}

func (c *Client) SendMessageToGroup(message message.MessageInf, group community.Group) error {
	// todo use go template
	//TODO implement me
	log.Infof("message: title: %v, body: %v, url: %v\n", message.GetTitle(), message.GetBody(), message.GetURL())
	return nil
}

func (c *Client) SendMessageToUsersInGroup(message message.MessageInf, users []community.User, group community.Group) error {
	// todo use go template
	log.Infof("SendMessageToUsers: %v\n", users[0].GetUserID()) // todo replace use slice with one user
	log.Infof("【title】: %v\n【body】: %v\n【url】: %v\n\n", message.GetTitle(), message.GetBody(), message.GetURL())
	return nil
}

func (c *Client) SendMessageToUserInGroup(message message.MessageInf, user community.User, group community.Group) error {
	// todo use go template
	log.Infof("SendMessageToUsers: %v\n", user.GetUserID())
	log.Infof("【title】: %v\n【body】: %v\n【url】: %v\n\n", message.GetTitle(), message.GetBody(), message.GetURL())
	return nil
}

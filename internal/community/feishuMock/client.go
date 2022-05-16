package feishuMock

import (
	"fmt"
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
	fmt.Printf("SendMessageToUsers: %v\n", users)
	fmt.Printf("message: title: %v, body: %v, url: %v\n", message.GetTitle(), message.GetBody(), message.GetURL())
	return nil
}

func (c *Client) SendMessageToGroup(message message.MessageInf, group community.Group) error {
	// todo use go template
	//TODO implement me
	fmt.Printf("message: title: %v, body: %v, url: %v\n", message.GetTitle(), message.GetBody(), message.GetURL())
	return nil
}

func (c *Client) SendMessageToUsersInGroup(message message.MessageInf, users []community.User, group community.Group) error {
	// todo use go template
	fmt.Printf("SendMessageToUsers: %v\n", users[0].GetUserID()) // todo replace use slice with one user
	fmt.Printf("【title】: %v\n【body】: %v\n【url】: %v\n\n", message.GetTitle(), message.GetBody(), message.GetURL())
	return nil
}

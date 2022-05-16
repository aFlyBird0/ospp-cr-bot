package community

import (
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/message"
)

// define abstract community, such as feishu(lark), dingtalk

type (
	Type string // community type
)

type User interface {
	GetUserID() string
	OfCommunity() Community
}

// Community is used to send message to reviewers, such as lark, dingtalk
type Community interface {
	GetType() Type
	GetUserByID(userID string) (User, error)
	GetGroupByID(groupID string) (Group, error)
	//MessageFromIssue(issue git.Issue) message.Message
	//MessageFromPr(pr git.PullRequest) message.Message
	//MessageFromComment(comment git.Comment) message.Message
	SendMessageToUsers(message message.MessageInf, users []User) error
	SendMessageToGroup(message message.MessageInf, group Group) error
	SendMessageToUsersInGroup(message message.MessageInf, users []User, group Group) error
	SendMessageToUserInGroup(message message.MessageInf, users User, group Group) error
}

type NeedRefresh interface {
	Refresh() error
}

// Group is a group of certain community
type Group interface {
	GetGroupID() string
	OfCommunity() Community
}

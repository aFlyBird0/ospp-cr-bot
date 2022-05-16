package feishu

import (
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/message"
)

type Message struct {
}

func (m Message) GetURL() string {
	//TODO implement me
	panic("implement me")
}

func (m Message) GetType() message.Type {
	//TODO implement me
	panic("implement me")
}

func (m Message) GetTitle() string {
	//TODO implement me
	panic("implement me")
}

func (m Message) GetBody() string {
	//TODO implement me
	panic("implement me")
}

func (m Message) GetToGitUser() git.User {
	//TODO implement me
	panic("implement me")
}

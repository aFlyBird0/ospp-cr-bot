package message

import "github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"

// define abstract messages

type Type string

// todo in progress
type MessageInf interface {
	GetURL() string
	GetType() Type
	GetTitle() string
	GetBody() string
	GetToGitUser() git.User
}

type Message struct {
	ToGitUser git.User
	URL       string
	Type      Type
	Title     string
	Body      string
	//Metadata  map[string]string
}

func (m *Message) GetToGitUser() git.User {
	return m.ToGitUser
}

func (m *Message) GetURL() string {
	return m.URL
}

func (m *Message) GetType() Type {
	return m.Type
}

func (m *Message) GetTitle() string {
	return m.Title
}

func (m *Message) GetBody() string {
	return m.Body
}

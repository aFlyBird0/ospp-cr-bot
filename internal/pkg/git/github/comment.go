package github

import (
	"time"

	"github.com/google/go-github/v42/github"

	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"
)

type Comment struct {
	CreateTime time.Time
	UpdateTime time.Time

	ID   int64
	Body string
	User *User
	URL  string
}

// todo to distinguish different types of comments
func toCommentInf(comment *github.PullRequestComment) *Comment {
	return &Comment{
		CreateTime: comment.GetCreatedAt(),
		UpdateTime: comment.GetCreatedAt(),
		ID:         comment.GetID(),
		Body:       comment.GetURL(),
		User:       toUserInf(comment.GetUser()),
		URL:        comment.GetHTMLURL(),
	}
}

func (c Comment) CreatedAt() time.Time {
	return c.CreateTime
}

func (c Comment) UpdatedAt() time.Time {
	return c.UpdateTime
}

func (c Comment) GetID() int64 {
	return c.ID
}

func (c Comment) GetBody() string {
	return c.Body
}

func (c Comment) GetUser() git.User {
	return c.User
}

func (c Comment) GetURL() string {
	return c.URL
}

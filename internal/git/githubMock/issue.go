package githubMock

import (
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"
	"time"
)

type Issue struct {
	Number int
	Title  string
	Body   string
	Labels []string
}

func (i *Issue) CreatedAt() time.Time {
	return time.Now()
}

func (i *Issue) UpdatedAt() time.Time {
	return time.Now()
}

func (i *Issue) GetID() int64 {
	return 123
}

func (i *Issue) GetNumber() int {
	return i.Number
}

func (i *Issue) GetTitle() string {
	return i.Title
}

func (i *Issue) GetBody() string {
	return i.Body
}

func (i *Issue) GetLabels() []string {
	return i.Labels
}

func (i *Issue) GetAssignees() []string {
	return []string{}
}

func (i *Issue) ListComments() []git.Comment {
	return []git.Comment{}
}

func (i *Issue) GetState() git.IssueState {
	return git.IssueStateOpen
}

func (i *Issue) GetURL() string {
	return "https://example-issue.com"
}

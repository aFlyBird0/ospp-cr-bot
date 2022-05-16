package githubMock

import "github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"

type PR struct {
	Number    int
	Title     string
	Body      string
	Assignees []*User
}

func (pr PR) GetID() int64 {
	return 456
}

func (pr PR) GetNumber() int {
	return pr.Number
}

func (pr PR) GetState() git.PrState {
	return git.PrStateOpen
}

func (pr PR) GetTitle() string {
	return pr.Title
}

func (pr PR) GetBody() string {
	return pr.Body
}

func (pr PR) ListLabels() []string {
	return []string{"label1", "label2"}
}

func (pr PR) GetCommitter() git.User {
	return nil
}

func (pr PR) ListAssignees() []git.User {
	users := make([]git.User, len(pr.Assignees))
	for i, user := range pr.Assignees {
		users[i] = user
	}
	return users
}

func (pr PR) GetURL() string {
	return "https://example-pr.com"
}

func (pr PR) ListComments() []git.Comment {
	return nil
}

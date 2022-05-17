package github

import (
	"github.com/google/go-github/v42/github"

	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"
)

type PullRequest struct {
	ID     int64
	Number int
	Title  string
	Body   string
	//Merged    bool
	//Mergeable bool
	State  git.PrState
	URL    string
	Labels []string

	Committer *User
	Assignees []*User
	Comments  []*Comment
}

func toPrInf(pr *github.PullRequest) *PullRequest {
	ret := &PullRequest{
		ID:     pr.GetID(),
		Number: pr.GetNumber(),
		Title:  pr.GetTitle(),
		Body:   pr.GetBody(),
		State:  git.PrState(pr.GetState()),
		URL:    pr.GetHTMLURL(),
	}

	if pr.GetUser() != nil {
		ret.Committer = toUserInf(pr.GetUser())
	}

	if len(pr.Assignees) > 0 {
		ret.Assignees = make([]*User, len(pr.Assignees))
		for i, assignee := range pr.Assignees {
			ret.Assignees[i] = toUserInf(assignee)
		}
	}

	if len(pr.Labels) > 0 {
		ret.Labels = make([]string, len(pr.Labels))
		for i, label := range pr.Labels {
			ret.Labels[i] = label.GetName()
		}
	}

	// todo decide if comments are needed
	//if len(pr.Comments) > 0 {
	//	ret.Comments = make([]*Comment, len(pr.GetComments()))
	//	for i, comment := range pr.GetComments() {
	//		ret.Comments[i] = toCommentInf(comment)
	//	}
	//}

	return ret
}

func (p *PullRequest) GetID() int64 {
	return p.ID
}

func (p *PullRequest) GetNumber() int {
	return p.Number
}

func (p *PullRequest) GetState() git.PrState {
	return p.State

}

func (p *PullRequest) GetTitle() string {
	return p.Title
}

func (p *PullRequest) GetBody() string {
	return p.Body
}

func (p *PullRequest) ListLabels() []string {
	return p.Labels
}

func (p *PullRequest) GetCommitter() git.User {
	return p.Committer
}

func (p *PullRequest) ListAssignees() []git.User {
	var users []git.User
	for _, user := range p.Assignees {
		users = append(users, user)
	}
	return users
}

func (p *PullRequest) GetURL() string {
	return p.URL
}

func (p *PullRequest) ListComments() []git.Comment {
	var comments []git.Comment
	for _, comment := range p.Comments {
		comments = append(comments, comment)
	}
	return comments
}

func stringToPrState(state string) git.PrState {
	switch state {
	case "open":
		return git.PrStateOpen
	case "closed":
		return git.PrStateClosed
	//case "merged":
	//	return git.PrStateMerged
	case "all":
		return git.PrStateAll
	default:
		return git.PrStateAll
	}
}

func prStateToString(state git.PrState) string {
	switch state {
	case git.PrStateOpen:
		return "open"
	case git.PrStateClosed:
		return "closed"
	case git.PrStateAll:
		return "all"
	default:
		return "all"
	}
}

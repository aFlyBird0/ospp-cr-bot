package github

import (
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"
	"github.com/google/go-github/v42/github"
	"time"
)

type Issue struct {
	ID         int64
	Number     int
	Title      string
	URL        string
	Body       string
	Labels     []*github.Label
	State      git.IssueState
	CreateTime time.Time
	UpdateTime time.Time
	User       *User
	Assignees  []*User
	Comments   []*Comment
}

func toIssueInf(issue *github.Issue) *Issue {
	assignees := make([]*User, 0)
	for _, assignee := range issue.Assignees {
		assignees = append(assignees, toUserInf(assignee))
	}
	//todo
	//comments := make([]*Comment, 0)
	//for _, comment := range issue.Comments {
	//	comments = append(comments, toCommentInf(comment))
	//}
	return &Issue{
		ID:         issue.GetID(),
		Number:     issue.GetNumber(),
		Title:      issue.GetTitle(),
		URL:        issue.GetHTMLURL(),
		Body:       issue.GetBody(),
		Labels:     issue.Labels,
		State:      git.IssueState(issue.GetState()),
		CreateTime: issue.GetCreatedAt(),
		UpdateTime: issue.GetUpdatedAt(),
		User:       toUserInf(issue.GetUser()),
		Assignees:  assignees,
		Comments:   nil,
	}
}

func (i *Issue) CreatedAt() time.Time {
	return i.CreateTime
}

func (i *Issue) UpdatedAt() time.Time {
	return i.UpdateTime
}

func (i *Issue) GetID() int64 {
	return i.ID
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
	labels := make([]string, len(i.Labels))
	for i, label := range i.Labels {
		labels[i] = *label.Name
	}
	return labels
}

func (i *Issue) GetAssignees() []git.User {
	var assignees []git.User
	for _, assignee := range i.Assignees {
		assignees = append(assignees, assignee)
	}
	return assignees
}

func (i *Issue) ListComments() []git.Comment {
	var comments []git.Comment
	for _, comment := range i.Comments {
		comments = append(comments, comment)
	}
	return comments
}

func (i *Issue) GetState() git.IssueState {
	return i.State
}

func (i *Issue) GetURL() string {
	return i.URL
}

func stringToIssueState(state string) git.IssueState {
	switch state {
	case "open":
		return git.IssueStateOpen
	case "closed":
		return git.IssueStateClosed
	case "all":
		return git.IssueStateAll
	default:
		return git.IssueStateAll
	}
}

func issueStateToString(state git.IssueState) string {
	switch state {
	case git.IssueStateOpen:
		return "open"
	case git.IssueStateClosed:
		return "closed"
	case git.IssueStateAll:
		return "all"
	default:
		return "all"
	}
}

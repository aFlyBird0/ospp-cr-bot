package git

import "time"

// define abstract git platforms and concepts, such as  github, gitlab, etc.

type (
	PlatformType string
)

type TimeAt interface {
	CreatedAt() time.Time
	UpdatedAt() time.Time
}

type Issue interface {
	TimeAt

	GetID() int64
	GetNumber() int
	GetTitle() string
	GetBody() string
	GetLabels() []string
	GetAssignees() []string // todo replace with git.User
	ListComments() []Comment
	GetState() IssueState
	GetURL() string
}

type Comment interface {
	TimeAt

	GetID() int64
	GetBody() string
	GetUser() User
	GetURL() string
}

type IssueComment interface {
	Comment

	GetIssue() Issue
}

type Commit interface {
	TimeAt

	GetID() string
	GetMessage() string
	GetURL() string
	GetAuthor() User
	GetCommitter() User
}

type PullRequest interface {
	GetID() int64
	GetNumber() int
	GetState() PrState
	GetTitle() string
	GetBody() string
	ListLabels() []string
	GetCommitter() User
	ListAssignees() []User
	GetURL() string

	ListComments() []Comment
}

type User interface {
	GetUserID() string
	OfPlatform() Platform
}

type Repo interface {
	//GetID() int64
	GetName() string
	GetDescription() string
	GetURL() string
	OfPlatForm() Platform
}

type Platform interface {
	GetType() PlatformType
	GetRepoInfo(repoName string) (Repo, error)
	ListRepos() []Repo
	ListIssuesByRepo(repo Repo) ([]Issue, error)
	ListIssuesByRepoWithFilter(repo Repo, filter IssueFilter) ([]Issue, error)
	ListPrsByRepo(repo Repo) ([]PullRequest, error)
	ListPrsByRepoWithFilter(repo Repo, filter PrFilter) ([]PullRequest, error)
}

// IssueFilter is used to filter issues by state
// todo add more filter options
type IssueFilter struct {
	State IssueState
}

// PrFilter is used to filter pull requests
// todo add more filter options
type PrFilter struct {
	State PrState
}
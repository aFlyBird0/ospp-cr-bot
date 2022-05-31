package git

type (
	IssueState string
	PrState    string
)

const (
	IssueStateAll    IssueState = ""
	IssueStateOpen   IssueState = "open"
	IssueStateClosed IssueState = "closed"

	PrStateAll    PrState = ""
	PrStateOpen   PrState = "open"
	PrStateClosed PrState = "closed"
	PrStateMerged PrState = "merged" //todo merged may be not supported by github
)

//const (
//	PlatformGitlab PlatformType = "gitlab"
//	PlatformGitee  PlatformType = "gitee"
//)

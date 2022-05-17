package union

import (
	"fmt"

	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/git"
	"github.com/devstream-io/devstream/ospp-cr-bot/pkg/message"
)

// AnalysePrCausedByWho Analyze PR was blocked by whom, and build message
// todo refactor: do not return message directly, but build message in community client
func AnalysePrCausedByWho(pr git.PullRequest) []message.Message {
	// todo complete logic of this function
	var messages []message.Message
	for _, assignee := range pr.ListAssignees() { // it is not easy to assign others because I only have one account
		//for _, assignee := range []git.User{pr.GetCommitter()} { // return committer and just for test
		messages = append(messages, message.Message{
			ToGitUser: assignee,
			Type:      message.MessageTypeNeedReivew,
			Title:     message.MessageTypeNeedReivew.Chinese(),
			Body:      fmt.Sprintf("PR [%v](#%v) is  blocked by %v", pr.GetTitle(), pr.GetNumber(), assignee.GetUserID()),
			URL:       pr.GetURL(),
		})
	}
	return messages
}

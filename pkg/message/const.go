package message

const (
	MessageTypeNeedReivew        Type = "need_review"
	MessageTypeNeedReplyOrCommit      = "need_reply_or_commit"
)

func (m Type) Chinese() string {
	switch m {
	case MessageTypeNeedReivew:
		return "待 Review"
	case MessageTypeNeedReplyOrCommit:
		return "待回复评论或 Commit"
	default:
		return "未知"
	}
}

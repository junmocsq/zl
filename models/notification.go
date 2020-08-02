package models

// 评论通知
type CommentNotification struct {
	Id            int
	Uid           int
	IsSend        int
	TopicId       int
	TopicTitle    string
	ParentId      int
	ParentUid     int
	ParentContent string
	Cid           int
	Cuid          int
	CContent      string
	CreateAt      int64
}

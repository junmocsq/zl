package models

// 主题评论
type TopicComment struct {
	Id        int
	Content   string `json:"content"  orm:"size(1000)"`
	Images    string `json:"images" description:"评论图片" orm:"size(150)"` // 子评论不能有图片
	Uid       int    `json:"uid"`
	Tid       int    `json:"tid"`
	Pid       int    `json:"pid"`
	ChildNum  int    `json:"child_num"`
	LikeNum   int    `json:"like_num"`
	UnlikeNum int    `json:"unlike_num"`
	IsDel     int8   `json:"is_del"`
	CreateAt  int64  `json:"create_at"`
}

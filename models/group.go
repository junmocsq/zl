package models

// 群组
type Group struct {
	Id       int
	Name     string `json:"name" orm:"size(10)"`
	Uid      string
	Desc     string `json:"desc" description:"群组说明" orm:"size(500)"`
	Avatar   string `json:"avatar" description:"群组头像" orm:"size(150)"`
	UserNum  int    `json:"user_num" description:"群人数"`
	TopicNum int    `json:"topic_num" description:"主题数"`
	IsDel    int8   `json:"is_del"`
	CreateAt int64  `json:"create_at"`
}

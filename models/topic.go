package models

// 主题
type Topic struct {
	Id        int    `json:"id"`
	Title     string `json:"title" orm:"size(200)"`
	Desc      string `json:"desc" description:"主题描述"  orm:"size(3000)"`
	Images    string `json:"images" orm:"size(1500)"`
	Gid       int    `json:"gid" description:"群组id"`
	Uid       int    `json:"uid"`
	LikeNum   int    `json:"like_num"`
	UnlikeNum int    `json:"unlike_num"`
	IsDel     int8   `json:"is_del"`
	CreateAt  int64  `json:"create_at"`
}

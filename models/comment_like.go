package models

type TopicCommentLike struct {
	Uid       int   `json:"uid"`
	Cid       int   `json:"cid"`
	IsLike    int8  `json:"is_like"`
	UpdatedAt int64 `json:"updated_at"`
}

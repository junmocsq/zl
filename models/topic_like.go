package models

type TopicLike struct {
	Uid      int   `json:"uid"`
	Tid      int   `json:"tid"`
	IsLike   int8  `json:"is_like"`
	UpdateAt int64 `json:"update_at"`
}

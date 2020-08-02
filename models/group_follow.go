package models

// 群组关注

type GroupFollow struct {
	Uid       int
	Gid       int
	IsManager int8  `json:"is_manager" description:"是否是管理人员 0 不是 1 是"`
	UpdateAt  int64 `json:"update_at"`
}

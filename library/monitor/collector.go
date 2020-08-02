package monitor

import (
	"time"
)

const MODULE_REQUEST = "request"

// 数据收集器
type Collector interface {
	Collect()
}

type RequestCollect struct {
	Key          string
	Module       string
	Params       map[string]string
	ResponseTime int64
	CreateTime   int64
}

func NewVisitCollect(key string, params map[string]string) *RequestCollect {
	r := &RequestCollect{}
	r.Module = MODULE_REQUEST
	r.Key = key
	r.Params = params
	r.CreateTime = time.Now().UnixNano()
}

func (r *RequestCollect) Collect() {
	now := time.Now().UnixNano()
	r.ResponseTime = now-r.CreateTime
	var p Parser = NewRequest()
	p.Add()
}

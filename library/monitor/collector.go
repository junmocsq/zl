package monitor

import (
	"time"
)

const MODULE_REQUEST = "request"
const COLLECT_CHANNEL_LEN = 10000

var collectChan = make(chan *Collect, 10000)

// 数据收集器
type Collector interface {
	Push()
}

type Collect struct {
	Key          string
	Module       string
	Params       map[string]string
	ResponseTime int64
	CreateTime   int64
}

func NewCollect(module, key string, params map[string]string) *Collect {
	r := &Collect{}
	r.Module = module
	r.Key = key
	r.Params = params
	r.CreateTime = time.Now().UnixNano()
	return r
}

func (r *Collect) Push() {
	now := time.Now().UnixNano()
	r.ResponseTime = now - r.CreateTime
	collectChan <- r
}

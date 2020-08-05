package monitor

import (
	"github.com/sirupsen/logrus"
)

func Stroage() {
	done := make(chan struct{})
	for i := 0; i < 20; i++ {
		go func() {
			p := Parse(done)
			for i := range p {
				// 存储
				logrus.WithField("monitor", "storage").
					Infof("createTime:%d responseTime:%dμs %s %s", i.CreateTime/1e9, i.ResponseTime/1e3, i.Key, i.Module)
			}
		}()
	}
}

package library

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func init() {
	RegisterConfig(os.Getenv("HOME") + "/www/wangqingshui")
}
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	c.SetEnv("dev")
	Convey("test Get String Int", t, func() {
		So("127.0.0.1", ShouldEqual, c.GetString("redis.host"))

		So(6379, ShouldEqual, c.GetInt("redis.port"))
	})

}

package redis

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCacheRedis(t *testing.T) {
	c := &cacheRedis{}
	key := "testwqs"
	val := "wangqingshui"
	expire := 600
	Convey("cache redis Set", t, func() {
		res, err := c.Set(key, val)
		So(err, ShouldBeNil)
		So(res, ShouldBeTrue)
		r, _ := c.Get(key)
		So(r, ShouldEqual, val)
	})

	Convey("cache redis Get", t, func() {
		c.Set(key, val)
		res, err := c.Get(key)
		So(err, ShouldBeNil)
		So(res, ShouldEqual, val)
	})

	Convey("cache redis SetEx", t, func() {
		res, err := c.SetEx(key, val, expire)
		So(err, ShouldBeNil)
		So(res, ShouldBeTrue)
		r, _ := ttl(key)
		So(r, ShouldBeGreaterThan, expire-5)
	})

	Convey("cache redis Del", t, func() {
		c.Set(key, val)
		res, err := c.Delete(key)
		t.Log("del")
		So(err, ShouldBeNil)
		So(res, ShouldBeTrue)
	})

	Convey("cache redis SetTimeout", t, func() {
		c.Set(key, val)
		res, err := c.SetTimeout(key, expire)
		So(err, ShouldBeNil)
		So(res, ShouldBeTrue)
		r, _ := ttl(key)
		So(r, ShouldBeGreaterThan, expire-5)
	})

}

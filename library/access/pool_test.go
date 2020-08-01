package access

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAccessor(t *testing.T) {
	access := NewAccessor()
	Convey("accessor pool", t, func() {
		length := access.Length()
		a, err := access.Get()
		So(err, ShouldBeNil)
		So(access.Length(), ShouldEqual, length-1)
		access.Close(a)
		So(access.Length(), ShouldEqual, length)
	})
}

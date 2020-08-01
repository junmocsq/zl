package access

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRegisterDb(t *testing.T) {
	RegisterMasterDb(MYSQL_WQS, "mysql_goweb")
	RegisterSlaveDbSameMasterDb(MYSQL_WQS)
	Convey("获取主从库", t, func() {
		So(masterDb(MYSQL_WQS), ShouldNotBeNil)
		So(slaveDb(MYSQL_WQS), ShouldNotBeNil)
	})
}

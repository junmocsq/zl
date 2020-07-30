package library

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCreateUserToken(t *testing.T) {
	Convey("生成用户token",t, func() {
		token := CreateToken("user00000010")
		fmt.Println(token)
		So(token,ShouldNotBeEmpty)
	})
}
package too

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStr2bytes(t *testing.T) {
	Convey("TestStr2bytes", t, func() {
		result := Str2bytes("too")
		So(result, ShouldNotBeNil)
		So(len(result), ShouldEqual, 3)
		So(result, ShouldResemble, []byte{116, 111, 111})
	})
}
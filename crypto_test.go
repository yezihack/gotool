package too

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)



//
var text = strings.Repeat("v", 100)
func BenchmarkMd5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Md5(text)
	}
}
func BenchmarkMd5V(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Md5V(text)
	}
}
// BenchmarkMd5-8    	 2879212	       484 ns/op	     176 B/op	       3 allocs/op
// BenchmarkMd5V-8   	 2869645	       612 ns/op	      64 B/op	       2 allocs/op

func TestMd5V(t *testing.T) {
	Convey("If Md5", t, func() {
		So(Md5V("too"), ShouldEqual, "b403d3f0efbf4cb850d2d543758cb57c")
	})

}
func TestMd5(t *testing.T) {
	Convey("If Md5", t, func() {
		So(Md5("too"), ShouldEqual, "b403d3f0efbf4cb850d2d543758cb57c")
	})
}
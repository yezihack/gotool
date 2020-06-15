package too

import (
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
	
}
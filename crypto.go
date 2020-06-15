package too

import (
	"crypto/md5"
	"encoding/hex"
)

// 加密类

// str To md5
func Md5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
// str To md5, but more power
func Md5V(text string) string {
	hash := md5.Sum(Str2bytes(text))
	return hex.EncodeToString(hash[:])
}

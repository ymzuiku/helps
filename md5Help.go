package helps

import (
	"crypto/md5"
	"fmt"
)

// Md5 string to md5 string
func Md5(str string) string {
	bt := []byte(str + "workos-yan-01")
	has := md5.Sum(bt)
	return fmt.Sprintf("%x", has)
}

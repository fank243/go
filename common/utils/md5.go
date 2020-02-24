package utils

import (
	"crypto/md5"
	"fmt"
)

/**
md5 32位加密
@plainText 明文
@return 32位密文
*/
func Md5Encode(plainText string) string {
	data := []byte(plainText)
	return fmt.Sprintf("%x", md5.Sum(data))
}

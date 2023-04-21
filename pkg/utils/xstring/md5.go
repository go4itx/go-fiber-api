package xstring

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 ...
func MD5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

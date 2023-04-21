package xstring

import (
	"fmt"
	"regexp"
)

// Extract 通过正则获取字符串中的个字符
func Extract(str string, pattern string, index int) (res string, err error) {
	reg := regexp.MustCompile(pattern)
	match := reg.FindStringSubmatch(str)
	if len(match) > index {
		res = match[index]
	} else {
		err = fmt.Errorf("matching failed")
	}

	return
}

// Extract 通过正则获取字符串中的个字符
func ExtractAll(str string, pattern string) (res []string, err error) {
	reg := regexp.MustCompile(pattern)
	match := reg.FindStringSubmatch(str)
	if len(match) > 1 {
		res = match[1:]
	} else {
		err = fmt.Errorf("matching failed")
	}

	return
}

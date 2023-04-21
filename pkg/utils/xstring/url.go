package xstring

import (
	"log"
	"regexp"
	"strings"
)

// RemoveQueryString 移除url参数
func RemoveQueryString(url string) string {
	index := strings.Index(url, "?")
	if index != -1 {
		url = url[:index]
	}

	return url
}

// MatchURL 使用正则判断url是否合法
func MatchURL(url string, pattern string) bool {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		log.Println("Error compiling regular expression:", err)
		return false
	}

	return regex.MatchString(url)
}

// Http2https ...
func Http2https(url string) string {
	if strings.HasPrefix(url, "http://") {
		return "https://" + strings.TrimPrefix(url, "http://")
	}

	return url
}

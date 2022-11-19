package xstring

import (
	"strings"
	"unicode/utf8"
)

// Snake2Came 下划线转为驼峰格式
func Snake2came(str string, firstLetterCase bool) string {
	str = strings.TrimSpace(str)
	if utf8.RuneCountInString(str) < 2 {
		return str
	}

	var buff strings.Builder
	var temp string
	for k, r := range str {
		c := string(r)
		if c != "_" {
			if temp == "_" {
				c = strings.ToUpper(c)
			}

			if firstLetterCase && k == 0 {
				c = strings.ToUpper(c)
			}

			_, _ = buff.WriteString(c)
		}

		temp = c
	}

	return buff.String()
}

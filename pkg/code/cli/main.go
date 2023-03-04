package main

import (
	"fmt"
	"home/pkg/utils/xfile"
	"log"
	"os"
	"strings"
)

const (
	msgFile  = "/pkg/code/msg.go"
	codeFile = "/pkg/code/code.go"
)

func main() {
	dir, err := xfile.GetCurrentDirectory()
	if err != nil {
		log.Println("err", err)
		return
	}

	bytes, err := os.ReadFile(dir + codeFile)
	if err != nil {
		log.Println("err", err)
		return
	}

	statusMessage := ""
	arr := strings.Split(string(bytes), "\n")
	for _, v := range arr {
		str := strings.Trim(v, " ")
		if str != "" && strings.Contains(str, "@msg") {
			key := strings.Trim(str[:strings.Index(v, "=")], " ")
			val := strings.Trim(str[(strings.Index(v, "@msg")+4):], " ")
			statusMessage += fmt.Sprintf("%v:\"%v\",\n", key, val)
		}
	}

	msgTemplate := `package code

	// statusMessage NOTE: Keep this in sync with the status code list
	var statusMessage = map[int]string{
	{{data}}
	}`

	msgTemplate = strings.Replace(msgTemplate, "{{data}}", strings.TrimRight(statusMessage, "\n"), 1)
	err = os.WriteFile(dir+msgFile, []byte(msgTemplate), 0755)
	if err != nil {
		log.Println("err", err)
		return
	}

	fmt.Println("已生成对应错误码说明！")
}

package main

import (
	"flag"
	"fmt"
	"home/pkg/utils/xfile"
	"home/pkg/utils/xstring"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	TPL_PATH        = "cmd/cli/template/"
	CMD_MAIN        = "cmd.main.txt"
	CONTROLLER      = "controller.txt"
	CONTROLLER_INIT = "controller.init.txt"
	SERVICE         = "service.txt"
	SERVICE_INIT    = "service.init.txt"
)

var (
	Relation = map[string]string{
		CMD_MAIN:        "cmd/{project}/main.go",
		CONTROLLER:      "internal/{project}/controller/{project}.go",
		CONTROLLER_INIT: "internal/{project}/controller/init.go",
		SERVICE:         "internal/{project}/service/{project}.go",
		SERVICE_INIT:    "internal/{project}/service/init.go",
	}
)

type Project struct {
	Name      string
	Name2came string
}

func main() {
	name := flag.String("name", "", "project name")
	module := flag.String("module", "all", "enum: all/c/s")
	flag.Parse()
	if *name == "" {
		log.Println("Please enter the project name")
		return
	}

	modules := make([]string, 0)
	switch *module {
	case "all":
		for key := range Relation {
			modules = append(modules, key)
		}
	case "s":
		modules = []string{SERVICE}
	case "c":
		modules = []string{CONTROLLER}
	default:
		log.Println("unknown module")
		return
	}

	project := Project{
		Name:      *name,
		Name2came: xstring.Snake2came(*name, true),
	}

	for _, key := range modules {
		if err := gen(key, project); err != nil {
			log.Println(err)
			return
		}
	}
}

// gen 根据模板生成对应目标文件
func gen(tplName string, project Project) (err error) {
	tplFile := TPL_PATH + tplName
	if err = xfile.Exists(tplFile); err != nil {
		return
	}

	tpl, err := template.ParseFiles(tplFile)
	if err != nil {
		return
	}

	target, ok := Relation[tplName]
	if !ok {
		return fmt.Errorf("no %s in relation", tplName)
	}

	target = strings.ReplaceAll(target, "{project}", project.Name)
	if err = xfile.MakeDirectory(filepath.Dir(target)); err != nil {
		return
	}

	file, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return
	}

	return tpl.Execute(file, project)
}

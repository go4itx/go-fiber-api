package main

import (
	"errors"
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
	tplPath        = "cmd/cli/template/"
	cmdMain        = "cmd.main.txt"
	controller     = "controller.txt"
	controllerInit = "controller.init.txt"
	service        = "service.txt"
	serviceInit    = "service.init.txt"
)

var (
	Relation = map[string]string{
		cmdMain:        "cmd/{app}/main.go",
		controller:     "internal/{app}/controller/{targetName}.go",
		controllerInit: "internal/{app}/controller/init.go",
		service:        "internal/{app}/service/{targetName}.go",
		serviceInit:    "internal/{app}/service/init.go",
	}
)

type Project struct {
	App             string
	Target          string
	TargetName      string
	TargetName2came string
}

func main() {
	app := flag.String("app", "", "app name")
	target := flag.String("target", "all", "enum: all/c/s") // c->controller, s->service
	targetName := flag.String("targetName", "", "target name")
	flag.Parse()
	if *app == "" {
		log.Println("Please enter the app name")
		return
	}

	if *targetName == "" {
		log.Println("Please enter the target name")
		return
	}

	targets := make([]string, 0)
	switch *target {
	case "all":
		for key := range Relation {
			targets = append(targets, key)
		}
	case "s":
		targets = []string{service}
	case "c":
		targets = []string{controller}
	default:
		log.Println("unknown module")
		return
	}

	project := Project{
		App:             *app,
		Target:          *target,
		TargetName:      *targetName,
		TargetName2came: xstring.Snake2came(*targetName, true),
	}

	for _, key := range targets {
		if err := gen(key, project); err != nil {
			log.Println(err)
			return
		}
	}
}

// gen 根据模板生成对应目标文件
func gen(tplName string, project Project) (err error) {
	tplFile := tplPath + tplName
	if !xfile.Exists(tplFile) {
		return errors.New(tplFile + " does not exist!")
	}

	tpl, err := template.ParseFiles(tplFile)
	if err != nil {
		return
	}

	targetFile, ok := Relation[tplName]
	if !ok {
		return fmt.Errorf("no %s in relation", tplName)
	}

	targetFile = strings.ReplaceAll(targetFile, "{app}", project.App)
	targetFile = strings.ReplaceAll(targetFile, "{targetName}", project.TargetName)
	if xfile.Exists(targetFile) {
		return errors.New(targetFile + " already exists, please delete it first!")
	}

	if err = xfile.MakeDirectory(filepath.Dir(targetFile)); err != nil {
		return
	}

	file, err := os.OpenFile(targetFile, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return
	}

	return tpl.Execute(file, project)
}

package conf

import (
	"flag"
	"fmt"
	"path"
	"strings"

	"github.com/spf13/viper"
)

var Viper *viper.Viper

// init viper
func init() {
	arg := flag.String("config", "config/config.toml", "config file path")
	flag.Parse()
	Viper = New(*arg)
}

// New Viper
func New(filePath string) *viper.Viper {
	ext := path.Ext(filePath)
	dir, name := path.Split(filePath)
	if dir == "" {
		dir = "./"
	}

	v := viper.New()
	v.AddConfigPath(dir)
	v.SetConfigType(strings.TrimLeft(ext, "."))
	v.SetConfigName(strings.ReplaceAll(name, ext, ""))
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Fatal error config file: %v \n", err.Error()))
	}

	return v
}

// Load specify configuration
func Load(name string, data interface{}) (err error) {
	return Viper.UnmarshalKey(name, data)
}

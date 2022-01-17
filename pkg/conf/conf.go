package conf

import (
	"flag"
	"fmt"
	"path"
	"strings"

	"github.com/spf13/viper"
)

var config *viper.Viper

// init viper
func init() {
	arg := flag.String("config", "config/config.toml", "config file path")
	flag.Parse()
	config = New(*arg)
}

// New Viper
func New(filePath string) *viper.Viper {
	ext := path.Ext(filePath)
	dir, name := path.Split(filePath)
	v := viper.New()
	v.AddConfigPath(dir)
	v.SetConfigType(strings.TrimLeft(ext, "."))
	v.SetConfigName(strings.TrimRight(name, ext))
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Fatal error config file: %v \n", err.Error()))
	}

	return v
}

// Load specify configuration
func Load(name string, data interface{}) (err error) {
	return config.UnmarshalKey(name, data)
}

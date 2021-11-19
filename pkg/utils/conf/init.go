package conf

import (
	"flag"
	"fmt"
	"path"
	"strings"

	"github.com/spf13/viper"
)

var v *viper.Viper

// init viper
func init() {
	arg := flag.String("config", "config/config.toml", "config file path")
	flag.Parse()
	v = New(*arg)
}

// New Viper
func New(filePath string) *viper.Viper {
	ext := path.Ext(filePath)
	dir, name := path.Split(filePath)
	vn := viper.New()
	vn.AddConfigPath(dir)
	vn.SetConfigType(strings.TrimLeft(ext, "."))
	vn.SetConfigName(strings.TrimRight(name, ext))
	if err := vn.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Fatal error config file: %v \n", err.Error()))
	}

	return vn
}

// Load specify configuration
func Load(name string, config interface{}, vv ...*viper.Viper) (err error) {
	if len(vv) > 0 {
		return vv[0].UnmarshalKey(name, config)
	}

	return v.UnmarshalKey(name, config)
}

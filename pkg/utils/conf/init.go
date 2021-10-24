package conf

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"path"
	"strings"
)

var mainViper *viper.Viper

// init viper
func init() {
	arg := flag.String("config", "config/config.toml", "config file path")
	flag.Parse()
	mainViper = New(*arg)
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
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	return v
}

// Load specify configuration
func Load(name string, config interface{}, v ...*viper.Viper) (err error) {
	if len(v) > 0 {
		return v[0].UnmarshalKey(name, config)
	}

	return mainViper.UnmarshalKey(name, config)
}

package conf

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"path"
	"strings"
)

var v *viper.Viper

// init  viper
func init() {
	config := flag.String("config", "config/config.toml", "config file path")
	flag.Parse()

	fileExt := path.Ext(*config)
	fileDir, fileName := path.Split(*config)
	v = viper.New()
	v.AddConfigPath(fileDir)
	v.SetConfigType(strings.TrimLeft(fileExt, "."))
	v.SetConfigName(strings.TrimRight(fileName, fileExt))
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

// Load specify configuration
func Load(name string, config interface{}) (err error) {
	return v.UnmarshalKey(name, config)
}

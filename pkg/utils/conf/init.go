package conf

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"path"
	"strings"
)

// init  viper
func init() {
	config := flag.String("config", "config/config.toml", "config file path")
	flag.Parse()

	fileExt := path.Ext(*config)
	fileDir, fileName := path.Split(*config)
	viper.AddConfigPath(fileDir)
	viper.SetConfigType(strings.TrimLeft(fileExt, "."))
	viper.SetConfigName(strings.TrimRight(fileName, fileExt))
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

// Load specify configuration
func Load(name string, config interface{}) (err error) {
	return viper.UnmarshalKey(name, config)
}

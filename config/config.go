package config

import (
	"flag"
	"home/pkg/utils/xviper"

	"github.com/spf13/viper"
)

var (
	Viper *viper.Viper
)

// init ...
func init() {
	// 1.argument() 2.embedFile()
	Viper = xviper.New(argument())
}

// Load specify configuration
func Load(name string, data interface{}) (err error) {
	return Viper.UnmarshalKey(name, data)
}

// 方式一：通过读取命令行参数，支持http，https网络读取
func argument() xviper.Config {
	arg := flag.String("config", "config/config.toml", "config file path")
	flag.Parse()
	return xviper.Config{
		URL: *arg,
	}
}

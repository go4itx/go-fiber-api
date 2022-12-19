package conf

import (
	"flag"
	"home/pkg/utils/xviper"

	"github.com/spf13/viper"
)

var (
	Viper *viper.Viper
)

// init viper
func init() {
	vc := xviper.Config{}
	// 方式一: 需自己自己实现config.Init() 返回参数Param
	// vc = config.Init()
	// 方式二：通过读取命令行参数
	arg := flag.String("config", "config/config.toml", "config file path")
	flag.Parse()
	vc.URL = *arg
	// 创建Viper
	Viper = xviper.New(vc)
}

// Load specify configuration
func Load(name string, data interface{}) (err error) {
	return Viper.UnmarshalKey(name, data)
}

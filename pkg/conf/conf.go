package conf

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
	Viper = xviper.New(config())
}

// 方式一：通过读取命令行参数，支持http，https网络读取
func config() xviper.Config {
	arg := flag.String("config", "config/config.toml", "config file path")
	flag.Parse()
	return xviper.Config{
		URL: *arg,
	}
}

// 方式二: 打包配置文件，通过embed设置xviper.Config，如:
// //go:embed config.toml
// var data []byte
// func config() xviper.Config {
// 	return xviper.Config{
// 		URL:  "",
// 		Data: data,
// 		Type: "toml",
// 	}
// }

// Load specify configuration
func Load(name string, data interface{}) (err error) {
	return Viper.UnmarshalKey(name, data)
}

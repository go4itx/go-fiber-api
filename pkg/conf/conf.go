package conf

import (
	_ "embed"
	"flag"
	"home/pkg/utils/xviper"

	"github.com/spf13/viper"
)

var (
	Viper *viper.Viper
)

// //go:embed config-dev.toml
// var data []byte

// // init 方式一: 打包配置文件，在当前目录放置配置文件，通过embed
// func init() {
// 	vc := xviper.Config{
// 		URL:  "",
// 		Data: data,
// 		Type: "toml",
// 	}

// 	// 创建Viper
// 	Viper = xviper.New(vc)
// }

// init 方式二：通过读取命令行参数，支持http，https网络读取
func init() {
	arg := flag.String("config", "config/config.toml", "config file path")
	flag.Parse()
	vc := xviper.Config{
		URL: *arg,
	}

	// 创建Viper
	Viper = xviper.New(vc)
}

// Load specify configuration
func Load(name string, data interface{}) (err error) {
	return Viper.UnmarshalKey(name, data)
}

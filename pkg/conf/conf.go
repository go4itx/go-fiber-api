package conf

import (
	_ "embed"
	"flag"
	"home/pkg/utils/xviper"

	"github.com/spf13/viper"
)

var (
	Viper *viper.Viper
	// //go:embed config-dev.toml
	// data []byte
)

// init viper
func init() {
	vc := xviper.Config{}
	// 方式一: 打包配置文件，在当前目录放置配置文件，通过embed
	// vc.Data = data
	// vc.Type = "toml"

	// 方式二：通过读取命令行参数，支持http，https网络读取
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

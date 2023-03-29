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

// 方式一: 打包配置文件，通过embed设置xviper.Config，自己实现config()
// 方式二：通过读取命令行参数，支持http，https网络读取

// init ...
func init() {
	Viper = xviper.New(config())
}

func config() xviper.Config {
	arg := flag.String("config", "config/config.toml", "config file path")
	flag.Parse()
	return xviper.Config{
		URL: *arg,
	}
}

// Load specify configuration
func Load(name string, data interface{}) (err error) {
	return Viper.UnmarshalKey(name, data)
}

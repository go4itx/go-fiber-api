package xviper

import (
	"bytes"
	"errors"
	"home/pkg/utils/client"
	"log"
	"path"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	URL  string // 配置文件路径：目前只支持本地文件
	Data []byte // 通过bytes.NewReader加载
	Type string // 通过bytes.NewReader时，设置配置类型
}

// New Viper
func New(config Config) *viper.Viper {
	var (
		err        error
		configName string
		configPath string
		configType string

		v = viper.New()
	)

	if config.URL != "" {
		var fileName string
		ext := path.Ext(config.URL)
		configPath, fileName = path.Split(config.URL)
		configType = strings.TrimLeft(ext, ".")
		configName = strings.ReplaceAll(fileName, ext, "")
	}

	// 通过网络获取(http/https)
	if strings.Contains(config.URL, "http://") || strings.Contains(config.URL, "https://") {
		data, err := client.Request(config.URL).Result()
		if err != nil {
			panic(err)
		}

		config.Data = data
		config.Type = configType
	}

	// 优先通过数据加载
	if len(config.Data) > 0 && config.Type != "" {
		v.SetConfigType(config.Type)
		err = v.ReadConfig(bytes.NewReader(config.Data))
	} else if config.URL != "" {
		if configPath == "" {
			configPath = "./"
		}

		v.AddConfigPath(configPath)
		v.SetConfigType(configType)
		v.SetConfigName(configName)
		err = v.ReadInConfig()
	} else {
		err = errors.New("config param error")
	}

	if err != nil {
		log.Printf("Fatal error config file: %v \n", err)
		panic(err)
	}

	return v
}

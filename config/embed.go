package config

import (
	_ "embed"
	"home/pkg/utils/xviper"
)

//go:embed config.toml
var Data []byte

// 方式二: 嵌入配置文件，通过embed设置xviper.Config，如:
func embedFile() xviper.Config {
	return xviper.Config{
		URL:  "",
		Data: Data,
		Type: "toml",
	}
}

package config

import (
	"github.com/cinling/cin"
)

// 获取默认配置
func GetApp() *cin.Config {
	config := cin.NewConfig()
	return config
}

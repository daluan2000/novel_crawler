package config_manager_interface

import (
	"github.com/fsnotify/fsnotify"
	"time"
)

// CreateConfigManagerT 创建ConfigManager对象的工厂函数的type，
// configType是配置文件类型，可选json，yaml，下层用的是 github.com/spf13/viper包，更多支持的文件类型可去它文档查看
// configPath是配置文件所在目录，configName是配置文件名（不包含后缀），程序会自动扫描configPath包括的所有目录
// 但程序只会添加一个文件并读取文件里的配置，所以一个 ConfigManager 对象只能读取一个配置文件
type CreateConfigManagerT func(configType string, configPath []string, configName string) (ConfigManagerInterface, error)

// ConfigManagerInterface 是ConfigManager的借口，定义了所有会用到的方法
type ConfigManagerInterface interface {

	// Get GetString GetBool ...... 根据key获取值
	// Get 如果值不存在，返回nil
	Get(key string) interface{}
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetFloat64(key string) float64
	GetDuration(key string) time.Duration
	GetStringSlice(key string) []string

	// AddConfigWatcher callback在文件发生变化时会被调用
	AddConfigWatcher(callback func(changeEvent fsnotify.Event))

	// Unmarshal 把整个配置文件映射为一个结构体
	Unmarshal(p interface{}) error
	// UnmarshalKey 把某个key映射为结构体
	UnmarshalKey(key string, p interface{}) error
}

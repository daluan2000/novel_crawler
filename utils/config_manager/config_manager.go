package config_manager

import (
	"errors"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	ci "novel_crawler/utils/config_manager/config_manager_interface"
	"strings"
	"sync"
	"time"
)

func CreateConfigManager(configType string, configPath []string, configName string) (ci.ConfigManagerInterface, error) {
	if !isTypeValid(configType) {
		return nil, errors.New("unknown config type : " + configType)
	}
	if len(configPath) == 0 || configName == "" {
		return nil, errors.New("configPath or configName can not be empty")
	}
	cm := &configManager{
		configType: configType,
		configPath: configPath,
		configName: configName,
		watchers:   make([]func(changeEvent fsnotify.Event), 0),
		vp:         viper.New(),
	}
	cm.vp.SetConfigFile(configType)
	for _, v := range configPath {
		cm.vp.AddConfigPath(v)
	}
	cm.vp.SetConfigName(configName)

	// 每次文件发生变化时，都清空缓存
	cm.AddConfigWatcher(func(in fsnotify.Event) {
		cm.clearCache()
	})
	if err := cm.vp.ReadInConfig(); err == nil {
		return cm, nil
	} else {
		return nil, err
	}
}

func isTypeValid(configType string) bool {
	var allTypes = map[string]bool{
		"yaml": true,
		"json": true,
		"toml": true,
	}
	return allTypes[strings.ToLower(configType)]
}

type configManager struct {
	configType string
	configPath []string
	configName string
	watchers   []func(changeEvent fsnotify.Event)
	cache      sync.Map
	vp         *viper.Viper
}

// isCache doCache getCache clearCache四个函数的左右分别是：
// 判断该键值是否已缓存 缓存该键值 获取缓存值 清除所有缓存
func (c *configManager) isCache(key string) bool {
	_, b := c.cache.Load(key)
	return b
}
func (c *configManager) doCache(key string, value interface{}) {
	c.cache.Store(key, value)
}
func (c *configManager) getCache(key string) interface{} {
	v, _ := c.cache.Load(key)
	return v
}
func (c *configManager) clearCache() {
	var newSyncMap sync.Map
	c.cache = newSyncMap
}

// Get GetString .....
// 下面这几个get函数功能大同小异，都是先查看键值是否已缓存，如果已缓存那就从缓存里读，如果没有那就从viper对象里读，读完之后写入缓存
func (c *configManager) Get(key string) interface{} {
	if c.isCache(key) {
		return c.getCache(key)
	} else {
		v := c.vp.Get(key)
		c.doCache(key, v)
		return v
	}
}
func (c *configManager) GetString(key string) string {
	if c.isCache(key) {
		return c.getCache(key).(string)
	} else {
		v := c.vp.GetString(key)
		c.doCache(key, v)
		return v
	}
}
func (c *configManager) GetBool(key string) bool {
	if c.isCache(key) {
		return c.getCache(key).(bool)
	} else {
		v := c.vp.GetBool(key)
		c.doCache(key, v)
		return v
	}
}
func (c *configManager) GetInt(key string) int {
	if c.isCache(key) {
		return c.getCache(key).(int)
	} else {
		v := c.vp.GetInt(key)
		c.doCache(key, v)
		return v
	}
}
func (c *configManager) GetInt32(key string) int32 {
	if c.isCache(key) {
		return c.getCache(key).(int32)
	} else {
		v := c.vp.GetInt32(key)
		c.doCache(key, v)
		return v
	}
}
func (c *configManager) GetInt64(key string) int64 {
	if c.isCache(key) {
		return c.getCache(key).(int64)
	} else {
		v := c.vp.GetInt64(key)
		c.doCache(key, v)
		return v
	}
}
func (c *configManager) GetFloat64(key string) float64 {
	if c.isCache(key) {
		return c.getCache(key).(float64)
	} else {
		v := c.vp.GetFloat64(key)
		c.doCache(key, v)
		return v
	}
}
func (c *configManager) GetDuration(key string) time.Duration {
	if c.isCache(key) {
		return c.getCache(key).(time.Duration)
	} else {
		v := c.vp.GetDuration(key)
		c.doCache(key, v)
		return v
	}
}
func (c *configManager) GetStringSlice(key string) []string {
	if c.isCache(key) {
		return c.getCache(key).([]string)
	} else {
		v := c.vp.GetStringSlice(key)
		c.doCache(key, v)
		return v
	}
}

func (c *configManager) Unmarshal(p interface{}) error {
	return c.vp.Unmarshal(p)
}
func (c *configManager) UnmarshalKey(key string, p interface{}) error {
	return c.vp.UnmarshalKey(key, p)
}

/*
// 克隆方法有问题，有待修改
func (c *configManager) Clone(configType string, configPath []string, configName string) (config_manager_interface.ConfigManagerInterface, error) {
	if configName == "" {
		configName = c.configName
	}
	if configType == "" {
		configType = c.configType
	}
	if newCm, err := CreateConfigManager(c.configType, append(c.configPath, configPath...), configName); err == nil {
		for _, v := range c.watchers {
			newCm.AddConfigWatcher(v)
		}
		return newCm, nil
	} else {
		return nil, err
	}
}
*/

func (c *configManager) AddConfigWatcher(callback func(changeEvent fsnotify.Event)) {
	c.watchers = append(c.watchers, callback)
	c.vp.OnConfigChange(func(in fsnotify.Event) {
		for _, v := range c.watchers {
			v(in)
		}
	})
	c.vp.WatchConfig()
}

package info

import (
	u "net/url"
	"novel_crawler/crawler/info/info_interf"
)

type store struct {
}

func (g *store) GetInfo(url *u.URL) info_interf.Info {
	return infoMap[url.Hostname()]
}

// CreateStore 这个比较简单，就不建工厂接口了
func CreateStore() info_interf.InfoStore {
	return &store{}
}
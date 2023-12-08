package info

import (
	u "net/url"
	"novel_crawler/crawler/info/info_interf"
)

var Getter info_interf.InfoGetter = &getter{}

type getter struct {
}

func (g *getter) GetInfo(url *u.URL) info_interf.Info {
	return infoMap[url.Hostname()]
}

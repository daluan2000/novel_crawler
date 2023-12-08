package getter_info

import (
	u "net/url"
	"novel_crawler/crawler/getter_info/getter_info_interf"
)

type getter struct {
}

func (g *getter) GetInfo(url *u.URL) getter_info_interf.Info {
	return infoMap[url.Hostname()]
}

// CreateGetter 这个比较简单，就不建工厂接口了
func CreateGetter() getter_info_interf.InfoGetter {
	return &getter{}
}

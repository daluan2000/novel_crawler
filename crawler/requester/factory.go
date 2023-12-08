package requester

import (
	u "net/url"
	"novel_crawler/crawler/requester/requester_interf"
)

var Factory requester_interf.Factory = &factory{}

type factory struct {
}

func (f *factory) CreateRequester(url *u.URL) requester_interf.Requester {
	// 暂时只生产这一个类
	return CreateCommon(url)
}

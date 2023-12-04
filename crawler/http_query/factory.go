package http_query

import (
	u "net/url"
)

func CreateQuery(url *u.URL) Query {
	// 暂时只生产这一个类
	return CreateCommon(url)
}

package filter

import (
	u "net/url"
	"novel_crawler/crawler/filter/filter_interf"
)

var Factory filter_interf.Factory = &factory{}

type factory struct {
}

func (f *factory) CreateFilter(url *u.URL) filter_interf.Filter {
	return &common{}
}

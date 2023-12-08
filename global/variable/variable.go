package variable

import (
	"novel_crawler/crawler/filter/filter_interf"
	"novel_crawler/crawler/getter_info/getter_info_interf"
	"novel_crawler/crawler/getter_next/getter_next_interf"
	"novel_crawler/crawler/requester/requester_interf"
)

var (
	Requester requester_interf.Requester
)

var (
	GetterInfo            getter_info_interf.InfoGetter
	GetterNextContent     getter_next_interf.Getter
	GetterNextChapterList getter_next_interf.Getter
)

var (
	Filter filter_interf.Filter
)

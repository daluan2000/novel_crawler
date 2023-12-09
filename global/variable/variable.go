package variable

import (
	"novel_crawler/crawler/filter/filter_interf"
	"novel_crawler/crawler/getter_next/getter_next_interf"
	"novel_crawler/crawler/info/info_interf"
	"novel_crawler/crawler/requester/requester_interf"
)

var (
	Requester requester_interf.Requester
	InfoStore info_interf.InfoStore

	GetterNextContent     getter_next_interf.Getter
	GetterNextChapterList getter_next_interf.Getter

	Filter filter_interf.Filter
)

var (
	SaveTitle    = true
	RequestCount = 0
)

package variable

import (
	"novel_crawler/crawler/getter_next/getter_next_interf"
	"novel_crawler/crawler/requester/requester_interf"
)

var (
	Requester requester_interf.Requester
)

var (
	GetterNextContent     getter_next_interf.Getter
	GetterNextChapterList getter_next_interf.Getter
)

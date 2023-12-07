package filter_interf

import (
	u "net/url"
	"novel_crawler/crawler/chapter"
)

type Filter interface {
	Filter(chapters []chapter.Chapter) []chapter.Chapter
}

type Factory interface {
	CreateFilter(url *u.URL) Filter
}

package filter_interf

import (
	u "net/url"
	"novel_crawler/crawler/chapter/chapter_interf"
)

type Filter interface {
	Filter(chapters []chapter_interf.Chapter) []chapter_interf.Chapter
}

type Factory interface {
	CreateFilter(url *u.URL) Filter
}

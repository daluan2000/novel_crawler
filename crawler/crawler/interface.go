package crawler

import (
	u "net/url"
	"novel_crawler/crawler/chapter"
)

type Crawler interface {
	// FetchChapterList 获取章节列表
	FetchChapterList() ([]chapter.Chapter, error)
	// FetchChapterContent 获取某一章节内容
	FetchChapterContent(c *chapter.Chapter) error
	// GetUrl 获取url
	GetUrl() *u.URL
}

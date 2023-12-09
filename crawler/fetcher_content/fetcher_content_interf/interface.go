package fetcher_content_interf

import (
	u "net/url"
	"novel_crawler/crawler/chapter/chapter_interf"
)

type Fetcher interface {
	Fetch(c *chapter_interf.Chapter) error
}

type Factory interface {
	CreateFetcher(url *u.URL) Fetcher
}

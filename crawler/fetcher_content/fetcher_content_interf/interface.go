package fetcher_content_interf

import (
	u "net/url"
	"novel_crawler/crawler/chapter"
)

type Fetcher interface {
	Fetch(c *chapter.Chapter) error
}

type Factory interface {
	CreateFetcher(url *u.URL) Fetcher
}

package fetcher_list_interf

import (
	u "net/url"
	"novel_crawler/crawler/chapter/chapter_interf"
)

type Fetcher interface {
	Fetch(url *u.URL) ([]chapter_interf.Chapter, error)
}

type Factory interface {
	CreateFetcher(url *u.URL) Fetcher
}

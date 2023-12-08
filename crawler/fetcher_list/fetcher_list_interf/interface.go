package fetcher_list_interf

import (
	u "net/url"
	"novel_crawler/crawler/chapter"
)

type Fetcher interface {
	Fetch(url *u.URL) ([]chapter.Chapter, error)
}

type Factory interface {
	CreateFetcher(url *u.URL) Fetcher
}

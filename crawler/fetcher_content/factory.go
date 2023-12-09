package fetcher_content

import (
	u "net/url"
	"novel_crawler/crawler/fetcher_content/fetcher_content_interf"
	"novel_crawler/global/variable"
)

var Factory fetcher_content_interf.Factory = &factory{}

type factory struct {
}

func (f factory) CreateFetcher(url *u.URL) fetcher_content_interf.Fetcher {
	info := variable.InfoStore.GetInfo(url)

	if info.MultiPageContent {
		return &multiPageFetcher{}
	} else {
		return &singlePageFetcher{}
	}
}

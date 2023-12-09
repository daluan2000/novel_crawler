package fetcher_list

import (
	u "net/url"
	"novel_crawler/crawler/fetcher_list/fetcher_list_interf"
	"novel_crawler/global/variable"
)

var Fatory fetcher_list_interf.Factory = &factory{}

type factory struct {
}

func (f *factory) CreateFetcher(url *u.URL) fetcher_list_interf.Fetcher {
	info := variable.InfoStore.GetInfo(url)
	if info.MultiPageChapterList {
		return &multiPageFetcher{}
	} else {
		return &singlePageFetcher{}
	}
}

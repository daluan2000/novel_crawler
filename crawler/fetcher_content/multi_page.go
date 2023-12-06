package fetcher_content

import (
	"novel_crawler/crawler/chapter"
)

type MultiPageFetcher struct {
}

func (m *MultiPageFetcher) Fetch(c *chapter.Chapter) error {

	// 这里根据c.Url创建query和getter
	//var query requester_interf.Requester
	//var getter getter_next_interf.Getter

	return nil
}

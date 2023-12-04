package fetcher_list

import (
	"novel_crawler/crawler/chapter"
)

type Fetcher interface {
	Fetch(c []chapter.Chapter) error
}

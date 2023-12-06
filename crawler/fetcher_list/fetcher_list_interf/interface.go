package fetcher_list_interf

import (
	"novel_crawler/crawler/chapter"
)

type Fetcher interface {
	Fetch(c []chapter.Chapter) error
}

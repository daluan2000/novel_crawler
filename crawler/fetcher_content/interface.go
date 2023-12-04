package fetcher_content

import (
	"novel_crawler/crawler/chapter"
)

type Fetcher interface {
	Fetch(c *chapter.Chapter) error
}

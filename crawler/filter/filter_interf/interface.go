package filter_interf

import "novel_crawler/crawler/chapter"

type Filter interface {
	Filter(chapters []chapter.Chapter) []chapter.Chapter
}

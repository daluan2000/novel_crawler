package filter

import "novel_crawler/crawler/chapter"

type common struct {
}

func (c *common) Filter(chapters []chapter.Chapter) []chapter.Chapter {
	m := make(map[string]int)
	for _, i := range chapters {
		m[i.Url.String()]++
	}

	idx := 0
	for ; idx < len(chapters); idx++ {
		if m[chapters[idx].Url.String()] == 1 {
			break
		}
	}

	chapters = chapters[idx:]
	for i, _ := range chapters {
		chapters[i].Number = i + 1
	}
	return chapters
}

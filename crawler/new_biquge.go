package crawler

// NewBiQuGeCrawler 以新笔趣阁为代表的一系列网站的爬虫器 http://www.xbiqugeo.com/shu/6420/
// 网站特点是，小说目录和章节内容可能以分页的形式展示
type NewBiQuGeCrawler struct {
}

func (b *NewBiQuGeCrawler) FetchChapterList() ([]Chapter, error) {
	r := make([]Chapter, 0)
	return r, nil
}

func (b *NewBiQuGeCrawler) FetchChapterContent(c *Chapter) error {
	return nil
}

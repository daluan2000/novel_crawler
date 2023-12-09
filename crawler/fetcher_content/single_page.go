package fetcher_content

import (
	"novel_crawler/crawler/chapter/chapter_interf"
	"novel_crawler/crawler/utils/str_util"
	"novel_crawler/global/variable"
)

type singlePageFetcher struct {
}

func (s *singlePageFetcher) Fetch(c *chapter_interf.Chapter) error {

	// 发起http请求，获取网页内容并解析
	dom, err := variable.Requester.CreateGoQuery(c.Url)
	if err != nil {
		return err
	}

	// 获取章节content
	c.ContentHtml, err = dom.Find(variable.InfoStore.GetInfo(c.Url).ContentSelector).Html()
	if err != nil {
		return err
	}

	// 有些网站采用gbk编码，这里编码格式统一调整为utf8
	if bts, err := str_util.GbkToUtf8([]byte(c.ContentHtml)); err == nil {
		c.ContentHtml = string(bts)
	} else {
		return err
	}

	return nil
}

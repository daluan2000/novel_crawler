package fetcher_content

import (
	"novel_crawler/crawler/chapter"
	"novel_crawler/crawler/utils/str_util"
	"novel_crawler/global/variable"
)

type MultiPageFetcher struct {
}

func (m *MultiPageFetcher) Fetch(c *chapter.Chapter) error {

	// 发起http请求，获取网页内容并解析
	dom, err := variable.Requester.CreateGoQuery(c.Url)
	if err != nil {
		return err
	}

	info := variable.InfoStore.GetInfo(c.Url)

	for {
		// 获取章节content
		ct, err := dom.Find(info.ContentSelector).Html()
		if err != nil {
			return err
		}

		// 有些网站采用gbk编码，这里编码格式统一调整为utf8
		if bts, err := str_util.GbkToUtf8([]byte(ct)); err == nil {
			ct = string(bts)
		} else {
			return err
		}
		c.ContentHtml += ct

		// 与“下一页”相关的操作
		{
			nextUrl, err := variable.GetterNextContent.NextUrl(dom)
			if err != nil {
				return err
			}
			if nextUrl == nil {
				break
			}

			dom, err = variable.Requester.CreateGoQuery(c.Url.ResolveReference(nextUrl))
			if err != nil {
				return err
			}
		}
	}

	return nil

}

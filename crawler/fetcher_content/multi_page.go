package fetcher_content

import (
	"errors"
	"novel_crawler/crawler/chapter"
	"novel_crawler/crawler/getter_next/getter_next_interf"
	"novel_crawler/crawler/requester/requester_interf"
	"novel_crawler/my_global"
	"novel_crawler/utils"
	"strings"
)

type MultiPageFetcher struct {
}

func (m *MultiPageFetcher) Fetch(c *chapter.Chapter) error {

	// 这里根据c.Url创建query和getter
	var query requester_interf.Requester
	var getter getter_next_interf.Getter

	// 发起http请求，获取网页内容并解析
	dom, err := query.CreateGoQuery(c.Url)

	if err != nil {
		return err
	}

	info := my_global.NewBiQuGeInfoByHost[c.Url.Hostname()]

	for {
		// 获取章节content
		ct, err := dom.Find(info.ContentSelector).Html()
		if err != nil {
			return err
		}

		// 有些网站采用gbk编码，这里编码格式统一调整为utf8
		if bts, err := utils.GbkToUtf8([]byte(ct)); err == nil {
			ct = string(bts)
		} else {
			return err
		}
		c.Content += ct

		// 与“下一页”相关的操作
		{
			nextUrl, err := getter.NextUrl(dom)
			if err != nil {
				return err
			}
			if nextUrl == nil {
				break
			}

			dom, err = query.CreateGoQuery(c.Url.ResolveReference(nextUrl))
			if err != nil {
				return err
			}
		}
	}

	// 删除content文本中的某些标签
	for _, v := range info.RemoveSelector {
		c.Content, err = utils.RemoveHtmlElem(c.Content, v)
		if err != nil {
			return err
		}
	}
	// 对content字符串进行替换
	for k, v := range info.StrReplace {
		c.Content = strings.Replace(c.Content, k, v, -1)
	}
	if len(c.Content) == 0 {
		return errors.New("empty content of chapter: " + c.Title)
	}
	return nil
}

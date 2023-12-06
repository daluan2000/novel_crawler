package fetcher_content

import (
	"errors"
	"novel_crawler/crawler/chapter"
	"novel_crawler/crawler/requester/requester_interf"
	"novel_crawler/my_global"
	"novel_crawler/utils"
	"strings"
)

type SinglePageFetcher struct {
}

func (s *SinglePageFetcher) FetchChapterContent(c chapter.Chapter) error {

	// 这里根据c.Url创建query
	var query requester_interf.Requester

	// 发起http请求，获取网页内容并解析
	dom, err := query.CreateGoQuery(c.Url)
	if err != nil {
		return err
	}

	// 获取章节content
	c.Content, err = dom.Find(my_global.BiQuGeInfoByHost[c.Url.Hostname()].ContentSelector).Html()
	if err != nil {
		return err
	}

	// 有些网站采用gbk编码，这里编码格式统一调整为utf8
	if bts, err := utils.GbkToUtf8([]byte(c.Content)); err == nil {
		c.Content = string(bts)
	} else {
		return err
	}

	// 删除content文本中的某些标签
	for _, v := range my_global.BiQuGeInfoByHost[c.Url.Hostname()].RemoveSelector {
		c.Content, err = utils.RemoveHtmlElem(c.Content, v)
		if err != nil {
			return err
		}
	}
	// 对content字符串进行替换
	rp := my_global.BiQuGeInfoByHost[c.Url.Hostname()].StrReplace
	for k, v := range rp {
		c.Content = strings.Replace(c.Content, k, v, -1)
	}

	if len(c.Content) == 0 {
		return errors.New("empty content of chapter: " + c.Title)
	}
	return nil
}

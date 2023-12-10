package getter_next

import (
	"github.com/PuerkitoBio/goquery"
	u "net/url"
	"novel_crawler/crawler/utils/str_util"
	"novel_crawler/global/variable"
	"strings"
)

type CommonChapterList struct {
}

func nextUrl(dom *goquery.Document, selector, subStr string) (*u.URL, error) {
	nextA := dom.Find(selector)
	if nextA.Length() == 0 {
		return nil, nil
	}

	// nextA是下一页的a标签的元素，bytesA是a标签文本转换为utf-8编码的字写流
	bytesA, err := str_util.GbkToUtf8([]byte(nextA.Text()))
	if err != nil {
		return nil, err
	}

	// 验证a标签的href元素
	href, ok := nextA.Attr("href")
	if strings.Contains(string(bytesA), subStr) && ok {
		return dom.Url.Parse(href)
	}
	return nil, nil
}
func (c *CommonChapterList) NextUrl(dom *goquery.Document) (*u.URL, error) {

	selector := variable.InfoStore.GetInfo(dom.Url).ChapterListNextSelector

	subStr := variable.InfoStore.GetInfo(dom.Url).ChapterListNextStr

	return nextUrl(dom, selector, subStr)
}

type CommonContent struct {
}

func (c *CommonContent) NextUrl(dom *goquery.Document) (*u.URL, error) {

	selector := variable.InfoStore.GetInfo(dom.Url).ContentNextSelector

	subStr := variable.InfoStore.GetInfo(dom.Url).ContentNextStr

	return nextUrl(dom, selector, subStr)
}

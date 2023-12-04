package getter_next

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"novel_crawler/my_global"
	"novel_crawler/utils"
	"strings"
)

type CommonChapterList struct {
}

func nextUrl(dom *goquery.Document, selector, subStr string) (*url.URL, error) {
	nextA := dom.Find(selector)
	if nextA.Length() == 0 {
		return nil, nil
	}

	// nextA是下一页的a标签的元素，bytesA是a标签文本转换为utf-8编码的字写流
	bytesA, err := utils.GbkToUtf8([]byte(nextA.Text()))
	if err != nil {
		return nil, err
	}

	// 验证a标签的href元素
	href, ok := nextA.Attr("href")
	if strings.Contains(string(bytesA), subStr) && ok {
		return url.Parse(href)
	}
	return nil, nil
}
func (c *CommonChapterList) NextUrl(dom *goquery.Document) (*url.URL, error) {

	selector := my_global.NewBiQuGeInfoByHost[dom.Url.Hostname()].ChapterListNextSelector

	subStr := my_global.NewBiQuGeInfoByHost[dom.Url.Hostname()].ChapterListNextStr

	return nextUrl(dom, selector, subStr)
}

type CommonContent struct {
}

func (c *CommonContent) NextUrl(dom *goquery.Document) (*url.URL, error) {

	selector := my_global.NewBiQuGeInfoByHost[dom.Url.Hostname()].ContentNextSelector

	subStr := my_global.NewBiQuGeInfoByHost[dom.Url.Hostname()].ContentNextStr

	return nextUrl(dom, selector, subStr)
}

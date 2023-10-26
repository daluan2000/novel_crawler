package crawler

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
)

type nextGetterCommon struct {
}

func (n *nextGetterCommon) NextUrl(dom *goquery.Document, selector, subStr string) (*url.URL, error) {
	nextA := dom.Find(selector)
	if nextA.Length() == 0 {
		return nil, nil
	}

	// nextA是下一页的a标签的元素，bytesA是a标签文本转换为utf-8编码的字写流
	bytesA, err := GbkToUtf8([]byte(nextA.Text()))
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

type chapterFilterCommon struct {
}

// Filter 对爬取到的目录进行一些过滤操作
// 笔趣阁目录页面，头部有一些章节重复 大概是这种形式 [6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8]，应该把头三个元素去掉 https://www.52bqg.org/book_128955/
func (c *chapterFilterCommon) Filter(chapters []Chapter) []Chapter {

	m := make(map[string]int)
	for _, i := range chapters {
		m[i.UrlStr]++
	}

	idx := 0
	for ; idx < len(chapters); idx++ {
		if m[chapters[idx].UrlStr] == 1 {
			break
		}
	}

	chapters = chapters[idx:]
	for i, _ := range chapters {
		chapters[i].Number = i + 1
	}
	return chapters
}

package crawler

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
)

// BiQuGeCrawler 以笔趣阁为代表的一系列网站的爬虫器 https://www.52bqg.org/
// 这类网站的特点：目录页面和章节内容页面均不分页展示，可以很简单地爬取
type BiQuGeCrawler struct {
	novelUrl *url.URL
}

// filterChapter 对爬取到的目录进行一些过滤操作
// 笔趣阁目录页面，头部有一些章节重复 大概是这种形式 [6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8]，应该把头三个元素去掉 https://www.52bqg.org/book_128955/
func filterChapter(chapters []Chapter) []Chapter {

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

func (b *BiQuGeCrawler) FetchChapterList() ([]Chapter, error) {
	dom, err := createGoQuery(b.novelUrl.String())
	if err != nil {
		return nil, err
	}

	// 获取章节目录信息
	r := make([]Chapter, 0)
	dom.Find(hostASelector[b.novelUrl.Hostname()]).Each(func(i int, selection *goquery.Selection) {
		// 获取a标签链接
		if path, ok := selection.Attr("href"); ok {
			// 把a标签链接转为url
			if pathUrl, err := url.Parse(path); err == nil {
				// 获取a标签文本，也就是标题内容
				if bts, err := GbkToUtf8([]byte(selection.Text())); err == nil {
					// 把获取到的信息push到r里面
					r = append(r, Chapter{
						UrlStr: b.novelUrl.ResolveReference(pathUrl).String(),
						Title:  string(bts),
					})
				}
			}
		}

	})

	r = filterChapter(r)
	return r, nil
}

func (b *BiQuGeCrawler) FetchChapterContent(c *Chapter) error {
	dom, err := createGoQuery(c.UrlStr)
	if err != nil {
		return err
	}

	c.Content, err = dom.Find(hostContentSelector[b.novelUrl.Hostname()]).Html()
	if err != nil {
		return err
	}

	if bts, err := GbkToUtf8([]byte(c.Content)); err == nil {
		c.Content = string(bts)
	} else {
		return err
	}

	rp := hostReplace[b.novelUrl.Hostname()]
	for i := 0; i < len(rp); i += 2 {
		c.Content = strings.Replace(c.Content, rp[i], rp[i+1], -1)
	}

	return nil
}

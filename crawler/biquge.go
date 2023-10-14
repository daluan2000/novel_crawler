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
func (b *BiQuGeCrawler) chapterFilter(chapters []Chapter) []Chapter {

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
	// 发起http请求，获取网页内容并解析
	dom, err := CreateGoQuery(b.novelUrl.String())
	if err != nil {
		return nil, err
	}

	// 获取章节目录信息
	r := make([]Chapter, 0)
	dom.Find(BiQuGeInfoByHost[b.novelUrl.Hostname()].ASelector).Each(func(i int, selection *goquery.Selection) {
		// 获取a标签链接
		if path, ok := selection.Attr("href"); ok {
			// 把a标签链接转为url
			if pathUrl, err := url.Parse(path); err == nil {
				// 获取a标签文本，也就是标题内容，有些网站采用gbk编码，这里编码格式统一调整为utf8
				if bts, err := GbkToUtf8([]byte(selection.Text())); err == nil {
					// 把获取到的信息append到r里面
					r = append(r, Chapter{
						UrlStr: b.novelUrl.ResolveReference(pathUrl).String(),
						Title:  string(bts),
					})
				}
			}
		}

	})

	r = b.chapterFilter(r)
	return r, nil
}

func (b *BiQuGeCrawler) FetchChapterContent(c *Chapter) error {
	// 发起http请求，获取网页内容并解析
	dom, err := CreateGoQuery(c.UrlStr)
	if err != nil {
		return err
	}

	// 获取章节content
	c.Content, err = dom.Find(BiQuGeInfoByHost[b.novelUrl.Hostname()].ContentSelector).Html()
	if err != nil {
		return err
	}

	// 有些网站采用gbk编码，这里编码格式统一调整为utf8
	if bts, err := GbkToUtf8([]byte(c.Content)); err == nil {
		c.Content = string(bts)
	} else {
		return err
	}

	// 删除content文本中的某些标签
	for _, v := range BiQuGeInfoByHost[b.novelUrl.Hostname()].RemoveSelector {
		c.Content, err = RemoveHtmlElem(c.Content, v)
		if err != nil {
			return err
		}
	}
	// 对content字符串进行替换
	rp := BiQuGeInfoByHost[b.novelUrl.Hostname()].StrReplace
	for k, v := range rp {
		c.Content = strings.Replace(c.Content, k, v, -1)
	}

	// 调试的时候用，把整个页面的html文本保存起来
	//{
	//	htmlText, err := dom.Html()
	//	if err != nil {
	//		return err
	//	}
	//	if bts, err := GbkToUtf8([]byte(htmlText)); err == nil {
	//		c.Content = "\n" + string(bts) + "\n"
	//	} else {
	//		return err
	//	}
	//}

	return nil
}

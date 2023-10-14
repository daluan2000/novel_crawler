package crawler

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/url"
	"strings"
)

// NewBiQuGeCrawler 以新笔趣阁为代表的一系列网站的爬虫器 http://www.xbiqugeo.com/shu/6420/
// 网站特点是，小说目录和章节内容可能以分页的形式展示
type NewBiQuGeCrawler struct {
	novelUrl *url.URL
}

func (n *NewBiQuGeCrawler) getNext(dom *goquery.Document, selector, subStr string) (*url.URL, error) {
	nextA := dom.Find(selector)
	if nextA.Length() == 0 {
		return nil, nil
	}

	// nextA是下一页的a标签的元素，bytesA是a标签文本转换为utf-8编码的字写流
	bytesA, err := GbkToUtf8([]byte(nextA.Text()))
	if err != nil {
		return nil, err
	}

	href, ok := nextA.Attr("href")
	if strings.Contains(string(bytesA), subStr) && ok {
		return url.Parse(href)
	}
	return nil, nil
}

func (n *NewBiQuGeCrawler) FetchChapterList() ([]Chapter, error) {
	log.Println("该网站章节是分页展示的，需要更长的时间爬取，大概需要几十秒的时间......")

	// 发起http请求，获取网页内容并解析
	dom, err := CreateGoQuery(n.novelUrl.String())
	if err != nil {
		return nil, err
	}

	r := make([]Chapter, 0)
	info := NewBiQuGeInfoByHost[n.novelUrl.Hostname()]

	for {
		// 把当前页包含的章节存起来
		dom.Find(info.ASelector).Each(func(i int, selection *goquery.Selection) {
			// 获取a标签链接
			if path, ok := selection.Attr("href"); ok {
				// 把a标签链接转为url
				if pathUrl, err := url.Parse(path); err == nil {
					// 获取a标签文本，也就是标题内容，有些网站采用gbk编码，这里编码格式统一调整为utf8
					if bts, err := GbkToUtf8([]byte(selection.Text())); err == nil {
						// 把获取到的信息append到r里面
						r = append(r, Chapter{
							UrlStr: n.novelUrl.ResolveReference(pathUrl).String(),
							Title:  string(bts),
						})
					}
				}
			}
		})

		// 与“下一页”相关的操作
		{
			nextUrl, err := n.getNext(dom, info.ChapterListNextSelector, info.ChapterListNextStr)
			if err != nil {
				return nil, err
			}
			if nextUrl == nil {
				break
			}

			dom, err = CreateGoQuery(n.novelUrl.ResolveReference(nextUrl).String())
			if err != nil {
				return nil, err
			}

		}

	}

	return r, nil
}

func (n *NewBiQuGeCrawler) FetchChapterContent(c *Chapter) error {

	// 发起http请求，获取网页内容并解析
	dom, err := CreateGoQuery(c.UrlStr)
	if err != nil {
		return err
	}

	info := NewBiQuGeInfoByHost[n.novelUrl.Hostname()]

	for {
		// 获取章节content
		ct, err := dom.Find(info.ContentSelector).Html()
		if err != nil {
			return err
		}

		// 有些网站采用gbk编码，这里编码格式统一调整为utf8
		if bts, err := GbkToUtf8([]byte(ct)); err == nil {
			ct = string(bts)
		} else {
			return err
		}
		c.Content += ct

		// 与“下一页”相关的操作
		{
			nextUrl, err := n.getNext(dom, info.ContentNextSelector, info.ContentNextStr)
			if err != nil {
				return err
			}
			if nextUrl == nil {
				break
			}

			dom, err = CreateGoQuery(n.novelUrl.ResolveReference(nextUrl).String())
			if err != nil {
				return err
			}
		}
	}

	// 删除content文本中的某些标签
	for _, v := range info.RemoveSelector {
		c.Content, err = RemoveHtmlElem(c.Content, v)
		if err != nil {
			return err
		}
	}
	// 对content字符串进行替换
	for k, v := range info.StrReplace {
		c.Content = strings.Replace(c.Content, k, v, -1)
	}

	return nil
}

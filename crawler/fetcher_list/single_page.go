package fetcher_list

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	u "net/url"
	"novel_crawler/crawler/chapter"
	"novel_crawler/crawler/utils/str_util"
	"novel_crawler/global/variable"
)

type SinglePageFetcher struct {
}

func (s *SinglePageFetcher) FetchChapterContent(url *u.URL) ([]chapter.Chapter, error) {

	// 发起http请求，获取网页内容并解析
	dom, err := variable.Requester.CreateGoQuery(url)
	if err != nil {
		return nil, err
	}

	// 获取章节目录信息
	r := make([]chapter.Chapter, 0)
	dom.Find(variable.InfoStore.GetInfo(url).ASelector).Each(func(i int, selection *goquery.Selection) {
		// 获取a标签链接
		if path, ok := selection.Attr("href"); ok {
			// 把a标签链接转为url
			if pathUrl, err := url.Parse(path); err == nil {
				// 获取a标签文本，也就是标题内容，有些网站采用gbk编码，这里编码格式统一调整为utf8
				if bts, err := str_util.GbkToUtf8([]byte(selection.Text())); err == nil {
					// 把获取到的信息append到r里面
					r = append(r, chapter.Chapter{
						Url:   pathUrl,
						Title: string(bts),
					})
				}
			}
		}

	})

	r = variable.Filter.Filter(r)
	if len(r) == 0 {
		return nil, errors.New("empty chapter list")
	}
	return r, nil
}

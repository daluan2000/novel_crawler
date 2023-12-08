package fetcher_list

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	u "net/url"
	"novel_crawler/crawler/chapter"
	"novel_crawler/crawler/utils/str_util"
	"novel_crawler/global/variable"
)

type MultiPageFetcher struct {
}

func (m *MultiPageFetcher) Fetch(url *u.URL) ([]chapter.Chapter, error) {

	// 发起http请求，获取网页内容并解析
	dom, err := variable.Requester.CreateGoQuery(url)
	if err != nil {
		return nil, err
	}

	r := make([]chapter.Chapter, 0)
	info := variable.GetterInfo.GetInfo(url)

	for {
		// 把当前页包含的章节存起来
		dom.Find(info.ASelector).Each(func(i int, selection *goquery.Selection) {
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

		// 与“下一页”相关的操作
		{

			nextUrl, err := variable.GetterNextChapterList.NextUrl(dom)
			if err != nil {
				return nil, err
			}
			if nextUrl == nil {
				break
			}

			dom, err = variable.Requester.CreateGoQuery(nextUrl)
			if err != nil {
				return nil, err
			}

		}
	}

	r = variable.Filter.Filter(r)
	if len(r) == 0 {
		return nil, errors.New("empty chapter list")
	}
	return r, nil

}

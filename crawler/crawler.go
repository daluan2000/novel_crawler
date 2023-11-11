package crawler

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"log"
	"net/http"
	u "net/url"
	"novel_crawler/global"
	"novel_crawler/utils"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

// Chapter 爬取流程相同的共用同一个实现类
type Chapter struct {
	Number  int
	UrlStr  string
	Title   string
	Content string
}

func (c *Chapter) Save(f *os.File) error {
	str := ""
	if global.SaveTitle {
		str = fmt.Sprintf("%s\n%s\n", c.Title, c.Content)
	} else {
		str = fmt.Sprintf("%s\n", c.Content)
	}
	_, err := f.WriteString(str)
	return err
}

type ChapterFilter interface {
	Filter(chapters []Chapter) []Chapter
}

type NextGetter interface {
	NextUrl(dom *goquery.Document, selector, subStr string) (*u.URL, error)
}

type CrawlerInterface interface {
	// FetchChapterList 获取章节列表
	FetchChapterList() ([]Chapter, error)
	// FetchChapterContent 获取某一章节内容
	FetchChapterContent(c *Chapter) error
	// GetUrl 获取url
	GetUrl() *u.URL
}

// Glc goroutine limit channel 限制并发量
// Gap 每一次请求的睡眠时间，限制吞吐量
var Glc = new(chan interface{})
var Gap = new(time.Duration)

// CreateGoQuery 所有的http请求都通过这里发送
func CreateGoQuery(urlStr string) (*goquery.Document, error) {

	var client = &http.Client{
		Timeout: time.Second * 15,
	}

	// 并发限制
	*Glc <- 1
	defer func() {
		if *Gap > 0 {
			time.Sleep(*Gap)
		}
		_ = <-*Glc
	}()

	req, _ := http.NewRequest("GET", urlStr, nil)
	req.Header.Set("User-Agent", utils.RandomUserAgent())

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// 别忘了释放链接
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error: " + err.Error())
		}
	}(resp.Body)
	if err != nil {
		return nil, err
	}

	dom, err := goquery.NewDocumentFromReader(resp.Body)
	return dom, err

}

// CreateCrawler 暂时只生产两个类
func CreateCrawler(novelUrlStr string) (CrawlerInterface, error) {

	novelUrl, err := u.Parse(novelUrlStr)
	if err != nil {
		return nil, err
	}
	if _, ok := global.BiQuGeInfoByHost[novelUrl.Hostname()]; ok {
		return &BiQuGeCrawler{
			novelUrl: novelUrl,
			filter:   &chapterFilterCommon{},
		}, nil
	}
	if _, ok := global.NewBiQuGeInfoByHost[novelUrl.Hostname()]; ok {
		return &NewBiQuGeCrawler{
			novelUrl:   novelUrl,
			nextGetter: &nextGetterCommon{},
		}, nil
	}
	return nil, errors.New("暂时不支持该网站")
}

// GbkToUtf8 GBK 转 UTF-8，如果本来就是UTF8那么本函数不进行任何操作
func GbkToUtf8(s []byte) ([]byte, error) {
	// 如果是uft8则直接返回
	if utf8.Valid(s) {
		return s, nil
	}
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func RemoveHtmlElem(str, selector string) (string, error) {

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(str))
	if err != nil {
		return "", err
	}

	// 删除符合seletor的元素
	dom.Find(selector).Remove()

	res, err := dom.Html()
	if err != nil {
		return "", err
	}

	res = res[25 : len(res)-14]
	return res, nil
}

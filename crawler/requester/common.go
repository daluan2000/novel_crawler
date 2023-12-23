package requester

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	u "net/url"
	"novel_crawler/crawler/utils/common_util"
	"novel_crawler/global/consts"
	"novel_crawler/global/variable"
	"time"
)

type common struct {
}

// CreateGoQuery 所有的http请求都通过这里发送 这里不再进行并发限制
func (c *common) CreateGoQuery(url *u.URL) (*goquery.Document, error) {

	urlStr := url.String()

	variable.RequestCount++

	var client = &http.Client{
		Timeout: time.Second * 15,
	}

	req, _ := http.NewRequest("GET", urlStr, nil)
	req.Header.Set("User-Agent", common_util.RandomUserAgent())

	var resp *http.Response

	// 发起请求
	err := common_util.Retry(func() error {
		var err1 error
		resp, err1 = client.Do(req)
		return err1
	}, consts.RetryCount)

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
	if err != nil {
		return nil, err
	}

	// 手动赋值
	dom.Url = url
	return dom, nil

}

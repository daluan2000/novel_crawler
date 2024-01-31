package requester

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	u "net/url"
	"novel_crawler/crawler/utils/common_util"
	"novel_crawler/global/variable"
)

type common struct {
	Client *http.Client
}

// CreateGoQuery 所有的http请求都通过这里发送 这里不再进行并发限制
func (c *common) CreateGoQuery(url *u.URL) (*goquery.Document, error) {

	// 记录请求次数
	variable.RequestCount++

	// 生成请求对象
	req, _ := http.NewRequest("GET", url.String(), nil)
	req.Header.Set("User-Agent", common_util.RandomUserAgent())

	// 返回值
	var resp *http.Response

	// 这里应不应该加retry，这是个问题
	// 发起请求
	err := common_util.Retry(func() error {
		var err1 error
		resp, err1 = c.Client.Do(req)
		return err1
	}, variable.RetryCount)

	if err != nil {
		return nil, err
	}

	// 关闭空闲连接链接和释放资源
	defer func() {
		c.Client.CloseIdleConnections()
		err = resp.Body.Close()
		if err != nil {
			log.Println("Error: " + err.Error())
		}
	}()

	// 记录cookie
	c.Client.Jar.SetCookies(url, resp.Cookies())

	// 发起请求
	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	// 手动赋值
	dom.Url = url

	return dom, nil

}

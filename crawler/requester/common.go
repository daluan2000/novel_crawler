package requester

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	u "net/url"
	"novel_crawler/crawler/utils/retry"
	"novel_crawler/global/consts"
	"novel_crawler/global/variable"
	"novel_crawler/my_global"
	"novel_crawler/utils"
	"time"
)

type Common struct {
	// Glc goroutine limit channel 限制并发量
	// Gap 每一次请求的睡眠时间，限制吞吐量
	Glc chan interface{}
	Gap time.Duration
}

// CreateGoQuery 所有的http请求都通过这里发送
func (c *Common) CreateGoQuery(url *u.URL) (*goquery.Document, error) {

	urlStr := url.String()

	my_global.RequestCount++

	var client = &http.Client{
		Timeout: time.Second * 15,
	}

	// 并发限制
	c.Glc <- 1
	defer func() {
		if c.Gap > 0 {
			time.Sleep(c.Gap)
		}
		_ = <-c.Glc
	}()

	req, _ := http.NewRequest("GET", urlStr, nil)
	req.Header.Set("User-Agent", utils.RandomUserAgent())

	var resp *http.Response

	// 发起请求
	err := retry.Retry(func() error {
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
	return dom, err

}

func CreateCommon(url *u.URL) *Common {
	res := &Common{}
	fl := variable.InfoStore.GetInfo(url).FrequencyLimit
	res.Glc, res.Gap = make(chan interface{}, fl.Concurrent), fl.Gap
	return res
}

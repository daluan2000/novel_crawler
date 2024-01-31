package requester

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	u "net/url"
	"novel_crawler/crawler/requester/requester_interf"
	"time"
)

var Factory requester_interf.Factory = &factory{}

type factory struct {
}

func (f *factory) CreateRequester(url *u.URL) requester_interf.Requester {
	// 暂时只生产这一个类
	res := &common{
		Client: &http.Client{
			Timeout: time.Second * 15,
		},
	}

	if res.Client.Jar == nil {
		var err error
		res.Client.Jar, err = cookiejar.New(nil)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	return res
}

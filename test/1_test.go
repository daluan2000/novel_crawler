package test

import (
	"fmt"
	"novel_crawler/crawler"
	"testing"
)

func Test1(t *testing.T) {

	if dom, err := crawler.CreateGoQuery("https://www.xbiqugeo.com/info/6420/5897186_2.html"); err == nil {
		fmt.Println(dom.Find("#next_url").Text())
		fmt.Println(dom.Find("a").Text())
	}
}

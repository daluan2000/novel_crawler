package test

import (
	"fmt"
	"novel_crawler/crawler"
	"testing"
)

func Test1(t *testing.T) {

	str := "<div class=\"posterror\"><a href=\"javascript:report();\" class=\"red\">章节错误,点此举报(免注册)</a>,举报后维护人员会在两分钟内校正章节内容,请耐心等待,并刷新页面。</div>\n                <a href=\"javascript:addbookcase('192619','假太监：再不死我就当皇帝了','616716','第一章 千古第一忠臣')\" class=\"btn-addbs\">天才一秒记住xbiqugeo.com，最快更新最新章节！</a>111"
	if newStr, err := crawler.RemoveHtmlElem(str, "div"); err == nil {
		fmt.Print(newStr)
	} else {
		fmt.Println(err.Error())
	}
}
package utils

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"math/rand"
	"strings"
	"unicode/utf8"
)

const (
	textBlack = iota + 30
	textRed
	textGreen
	textYellow
	textBlue
	textPurple
	textCyan
	textWhite
)

//func init() {
//	bts, err := os.ReadFile("../utils/user_agents")
//	if err != nil {
//		log.Fatal(err)
//	}
//	userAgents = strings.Split(string(bts), "\r\n")
//}

func Purple(str string) string {
	return textColor(textPurple, str)
}
func Yellow(str string) string {
	return textColor(textYellow, str)
}
func Red(str string) string {
	return textColor(textRed, str)
}

func Green(str string) string {
	return textColor(textGreen, str)
}

func textColor(color int, str string) string {
	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, str)
}

func RandomUserAgent() string {
	// 推荐使用
	idx := rand.Int() % len(userAgents)
	return userAgents[idx]
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

func HtmlToText(str string) (string, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(str))
	if err != nil {
		return "", err
	}

	res := dom.Text()

	return res, nil
}

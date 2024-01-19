package str_util

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"
)

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
func RemovePreBlank(s string) string {
	blanks := []string{" ", "\r", "\n", "\t", "\v", "\f"}
	i := 0
	for i < len(s) {
		if slices.Contains(blanks, s[i:i+1]) {
			i++
		} else {
			break
		}
	}
	return s[i:]
}
func RemoveSufBlank(s string) string {
	blanks := []string{" ", "\r", "\n", "\t", "\v", "\f"}
	i := len(s) - 1
	for i >= 0 {
		if slices.Contains(blanks, s[i:i+1]) {
			i--
		} else {
			break
		}
	}
	return s[0 : i+1]
}

func RemovePreSufBlank(s string) string {
	return RemoveSufBlank(RemovePreBlank(s))
}

// TagRegexp 返回匹配标签的正则字符串，反引号包裹的字符串不会被转义
func TagRegexp(tag string) [2]string {
	st := fmt.Sprintf(`<%s\s*>|<%s\s+.*?>`, tag, tag)
	ed := fmt.Sprintf(`</%s>`, tag)
	_ = regexp.MustCompile(st)
	_ = regexp.MustCompile(ed)
	return [2]string{st, ed}
}

package chapter_interf

import (
	u "net/url"
)

// Chapter 爬取流程相同的共用同一个实现类
type Chapter struct {
	Number       int // 序号
	Url          *u.URL
	Title        string
	ContentHtml  string   // html文本
	ContentText  string   // text文本
	ContentFinal []string // 处理后最终写入的文本，每一个字符串代表一行
	Err          error
}

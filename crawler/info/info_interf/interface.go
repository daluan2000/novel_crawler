package info_interf

import (
	u "net/url"
	"time"
)

type Info struct {
	ASelector       string
	ContentSelector string

	// 关于字符串替换，有几点要注意：
	// 1. 代码文件的编码格式为utf-8，小说文件的编码格式也应该为utf-8，这样才能保证替换的结果是正确的
	// 2. UTF-8是Unicode的一种实现方式，某些非utf-8编码的字符，在进行编码转换后可能出现乱码现象，这时直接百度搜索对应的unicode编码，然后替换即可
	//    比如对于GBK编码下的&nbsp空格字符，在转换为utf8后它的显示为 聽 \u807d，下面一行做了替换
	StrReplace map[string]string
	// 字符串中删除一些标签
	RemoveSelector []string

	FrequencyLimit
	NextChapterList
	NextContent
}
type FrequencyLimit struct {
	// 并发量限制
	Concurrent int
	// 每次请求后线程的休眠时间
	Gap time.Duration
}

type NextChapterList struct {
	HasNextChapterList      bool
	ChapterListNextSelector string
	ChapterListNextStr      string
}

type NextContent struct {
	HasNextContent      bool
	ContentNextSelector string
	ContentNextStr      string
}

type InfoStore interface {
	GetInfo(url *u.URL) Info
}

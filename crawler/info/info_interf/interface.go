package info_interf

import (
	u "net/url"
	"time"
)

// Info结构体
type Info struct {
	// 目录页面，各章节标题a标签的选择器
	ASelector string
	// 章节内容页面，小说内容的选择器
	ContentSelector string
	// html文本替换字符串
	StrReplace map[string]string
	// html文本替换字符串，正则表达式形式
	RegReplace map[string]string
	// html文本中要删除的标签对应的选择器
	RemoveSelector []string

	FrequencyLimit
	NextChapterList
	NextContent
}

// 并发限制，有默认值
type FrequencyLimit struct {
	// 并发数量限制
	Concurrent int
	// 每次请求后线程的休眠时间
	Gap time.Duration
}

// 目录为分页展示时，需要加上此部分信息
type NextChapterList struct {
	// 如果分页展示，设置为true
	MultiPageChapterList bool
	// 目录页面中，下一页a标签的选择器
	ChapterListNextSelector string
	// 目录页面中，下一页a标签应包含的文本
	ChapterListNextStr string
}

// 章节内容分页展示时，需要加上此部分信息
type NextContent struct {
	// 如果分页展示，设置为true
	MultiPageContent bool
	// 章节内容页面中，下一页a标签的选择器
	ContentNextSelector string
	// 章节内容页面中，下一页a标签应包含的文本
	ContentNextStr string
}

// SameInfo 记录有相同info的网站，基础info是相同的，但并发限制可以自定义
type SameInfo struct {
	Host string
	FrequencyLimit
}

type InfoStore interface {
	GetInfo(url *u.URL) Info
	ReadYaml(fileName string) error
	Exist(url *u.URL) bool
	FillInfoDefault()
}

package chapter

import (
	u "net/url"
	"novel_crawler/my_global"
	"os"
	"strings"
)

// Chapter 爬取流程相同的共用同一个实现类
type Chapter struct {
	Number       int // 序号
	Url          *u.URL
	Title        string
	ContentHtml  string   // html文本
	ContentText  string   // text文本
	ContentFinal []string // 处理后最终写入的文本，每一个字符串代表一行
}

func (c *Chapter) Save(f *os.File) error {

	str := ""
	if my_global.SaveTitle {
		str += c.Title + "\n"
	}
	str += "支持正版，人人有责\n支持正版，人人有责\n"
	str += strings.Join(c.ContentFinal, "\n")
	_, err := f.WriteString(str)
	return err
}

func (c *Chapter) GenerateText() error {
	// 删除html文本里的某些标签内容，然后替换掉其他标签，生成换行符
	return nil
}

func (c *Chapter) GenerateFinal() {

}

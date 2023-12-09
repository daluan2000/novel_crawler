package chapter

import (
	"errors"
	u "net/url"
	"novel_crawler/crawler/utils/str_util"
	"novel_crawler/global/variable"
	"novel_crawler/my_global"
	"novel_crawler/utils"
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
	Err          error
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
func (c *Chapter) DoBeforeSave() error {
	err := c.generateText()
	if err != nil {
		return err
	}
	c.generateFinal()
	return nil
}

func (c *Chapter) generateText() error {
	// 删除content文本中的某些标签
	var err error
	for _, v := range variable.InfoStore.GetInfo(c.Url).RemoveSelector {
		c.ContentText, err = utils.RemoveHtmlElem(c.ContentHtml, v)
		if err != nil {
			return err
		}
	}

	// 对text进行替换
	for k, v := range variable.InfoStore.GetInfo(c.Url).StrReplace {
		c.ContentText = strings.Replace(c.ContentText, k, v, -1)
	}

	if len(c.ContentText) == 0 {
		return errors.New("empty content of chapter: " + c.Title)
	}
	return nil
}

func (c *Chapter) generateFinal() {
	finalContent := strings.Split(c.ContentText, "\n")
	for i := 0; i < len(finalContent); i++ {
		finalContent[i] = str_util.RemovePreSufBlank(finalContent[i])
	}

	c.ContentFinal = make([]string, 0)
	for i := 0; i < len(finalContent); i++ {
		if finalContent[i] != "" {
			c.ContentFinal = append(c.ContentFinal, finalContent[i])
		}
	}

}

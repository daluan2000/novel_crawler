package chapter_handler

import (
	"errors"
	"novel_crawler/crawler/chapter/chapter_interf"
	"novel_crawler/crawler/utils/str_util"
	"novel_crawler/global/variable"
	"novel_crawler/my_global"
	"novel_crawler/utils"
	"os"
	"strings"
)

type Handler struct {
}

func (h *Handler) Save(f *os.File, c *chapter_interf.Chapter) error {
	str := ""
	if my_global.SaveTitle {
		str += c.Title + "\n"
	}
	str += "支持正版，人人有责\n支持正版，人人有责\n"
	str += strings.Join(c.ContentFinal, "\n") + "\n"
	_, err := f.WriteString(str)
	return err
}

func (h *Handler) generateText(c *chapter_interf.Chapter) error {
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

func (h *Handler) generateFinal(c *chapter_interf.Chapter) {
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
func (h *Handler) DoBeforeSave(c *chapter_interf.Chapter) error {
	err := h.generateText(c)
	if err != nil {
		return err
	}
	h.generateFinal(c)
	return nil
}

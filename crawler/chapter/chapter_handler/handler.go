package chapter_handler

import (
	"errors"
	"novel_crawler/crawler/chapter/chapter_interf"
	"novel_crawler/crawler/utils/str_util"
	"novel_crawler/global/variable"
	"os"
	"regexp"
	"strings"
)

type Handler struct {
}

func (h *Handler) Save(f *os.File, c *chapter_interf.Chapter) error {
	str := ""
	if variable.SaveTitle {
		str += c.Title + "\n"
	}
	str += "begin\n"
	str += strings.Join(c.ContentFinal, "\n") + "\n"
	_, err := f.WriteString(str)
	return err
}

func (h *Handler) generateTitle(c *chapter_interf.Chapter) error {
	c.Title = str_util.RemovePreSufBlank(c.Title)
	return nil
}

func (h *Handler) generateText(c *chapter_interf.Chapter) error {
	// 删除content文本中的某些标签
	var err error
	c.ContentText = c.ContentHtml
	for _, v := range variable.InfoStore.GetInfo(c.Url).RemoveSelector {
		c.ContentText, err = str_util.RemoveHtmlElem(c.ContentText, v)
		if err != nil {
			return err
		}
	}

	// 对text进行替换
	for k, v := range variable.InfoStore.GetBaseRegReplace() {
		reg := regexp.MustCompile(k)
		c.ContentText = reg.ReplaceAllString(c.ContentText, v)
	}
	for k, v := range variable.InfoStore.GetBaseStrReplace() {
		c.ContentText = strings.Replace(c.ContentText, k, v, -1)
	}
	for k, v := range variable.InfoStore.GetInfo(c.Url).StrReplace {
		c.ContentText = strings.Replace(c.ContentText, k, v, -1)
	}

	if len(c.ContentText) == 0 {
		return errors.New("empty content of chapter: " + c.Title)
	}
	return nil
}

func (h *Handler) generateFinal(c *chapter_interf.Chapter) error {
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
	return nil
}
func (h *Handler) DoBeforeSave(c *chapter_interf.Chapter) error {
	err := h.generateTitle(c)
	if err != nil {
		return err
	}
	err = h.generateText(c)
	if err != nil {
		return err
	}
	err = h.generateFinal(c)
	return err
}

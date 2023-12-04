package chapter

import (
	"fmt"
	u "net/url"
	"novel_crawler/my_global"
	"os"
)

// Chapter 爬取流程相同的共用同一个实现类
type Chapter struct {
	Number  int
	Url     *u.URL
	Title   string
	Content string
}

func (c *Chapter) Save(f *os.File) error {
	str := ""
	if my_global.SaveTitle {
		str = fmt.Sprintf("%s\n%s\n%s\n", c.Title, "    支持正版，人人有责", c.Content)
	} else {
		str = fmt.Sprintf("%s\n", c.Content)
	}
	_, err := f.WriteString(str)
	return err
}

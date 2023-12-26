package bootstrap

import (
	"log"
	u "net/url"
	"novel_crawler/crawler/chapter/chapter_handler"
	"novel_crawler/crawler/filter"
	"novel_crawler/crawler/getter_next"
	"novel_crawler/crawler/info"
	"novel_crawler/crawler/requester"
	"novel_crawler/crawler/utils/color_util"
	"novel_crawler/global/variable"
)

func init() {

	variable.InfoStore = info.CreateStore()
	if err := variable.InfoStore.ReadYaml("info"); err != nil {
		log.Println(color_util.Red("未发现配置文件或文件格式错误 " + err.Error()))
	} else {
		log.Println(color_util.Green("已读取配置文件"))
	}
	variable.InfoStore.FillInfoDefault()
}
func InitByUrl(url *u.URL) {
	variable.Requester = requester.Factory.CreateRequester(url)
	variable.Filter = filter.Factory.CreateFilter(url)
	variable.ChapterHandler = chapter_handler.Factory.CreateHandler(url)
	variable.GetterNextContent = getter_next.Factory.CreateContentNextGetter(url)
	variable.GetterNextChapterList = getter_next.Factory.CreateChapterListNextGetter(url)
}

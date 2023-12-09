package bootstrap

import (
	"log"
	u "net/url"
	"novel_crawler/crawler/filter"
	"novel_crawler/crawler/info"
	"novel_crawler/crawler/requester"
	"novel_crawler/crawler/utils/color_util"
	"novel_crawler/global/variable"
)

func init() {
	log.SetFlags(log.LstdFlags)

	if err := info.ReadYaml("info"); err != nil {
		log.Println(color_util.Red("未发现配置文件或文件格式错误 " + err.Error()))
	} else {
		log.Println(color_util.Green("已读取配置文件"))
	}

	info.InitInfo()
	variable.InfoStore = info.CreateStore()
}
func InitByUrl(url *u.URL) {
	variable.Requester = requester.Factory.CreateRequester(url)
	variable.Filter = filter.Factory.CreateFilter(url)
}

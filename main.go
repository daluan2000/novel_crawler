package main

import (
	"flag"
	"log"
	u "net/url"
	"novel_crawler/bootstrap"
	"novel_crawler/crawler/controller"
	"novel_crawler/crawler/utils/color_util"
	"novel_crawler/global/variable"
	"time"
)

func main() {

	var fileName = flag.String("f", "", "保存文件名")
	var urlStr = flag.String("u", "", "url链接")
	var saveTitle = flag.Int("st", 1, "保存tittle为1，不保存title为2，不输入该参数默认为1")
	flag.Parse()
	variable.SaveTitle = *saveTitle == 1

	url, err := u.Parse(*urlStr)
	if err != nil {
		log.Println(color_util.Red("错误的url格式"))
		return
	}

	bootstrap.InitByUrl(url)
	cl := controller.Factory.CreateController(url)
	cl.DoCrawling(url, *fileName+".txt")

	time.Sleep(time.Second)
}

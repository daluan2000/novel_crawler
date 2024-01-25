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

func parseFlag() {
	var fileName = flag.String("f", "", "保存文件名")
	var urlStr = flag.String("u", "", "url链接")
	var saveTitle = flag.Int("st", 1, "保存tittle为1，不保存title为2，不输入该参数默认为1")
	var logLevel = flag.Int("log", 1, "默认为1，打印详细log为2")
	var fillTitle = flag.Int("ft", 1, "不改变标题为1，填充标题编号为2，不输入该参数默认为1")
	var retryCount = flag.Int("rc", 10, "重新尝试的次数，默认为10")
	var retrySleep = flag.Duration("rs", 250*time.Millisecond, "retry时的休眠时间，默认250ms")
	flag.Parse()
	variable.FileName = *fileName
	variable.SaveTitle = *saveTitle == 1
	variable.FillTitle = *fillTitle == 2
	variable.RetryCount = *retryCount
	variable.RetrySleep = *retrySleep
	if *logLevel == 2 {
		log.SetFlags(log.LstdFlags | log.Llongfile)
	} else {
		log.SetFlags(log.Ltime)
	}

	url, err := u.Parse(*urlStr)
	if err != nil {
		log.Println(color_util.Red("错误的url格式"))
		return
	}

	variable.Url = url

}

func main() {

	parseFlag()

	bootstrap.InitByUrl(variable.Url)

	cl := controller.Factory.CreateController(variable.Url)
	cl.DoCrawling(variable.Url, variable.FileName+".txt")

	time.Sleep(time.Second)
}

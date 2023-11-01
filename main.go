package main

import (
	"flag"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
	"log"
	u "net/url"
	"novel_crawler/crawler"
	"novel_crawler/utils"
	"os"
	"time"
)

// 目前适配网站 https://www.52bqg.org/book_128955/

func retry(task func() error, count int) error {
	if err := task(); err == nil {
		return nil
	} else if count > 1 {
		//log.Println("do retry, remain retry count:", count-1)
		return retry(task, count-1)
	} else {
		return err
	}

}

func initConcurrentLimit(urlStr string) {
	glc := make(chan interface{}, 50)
	gap := time.Millisecond * 0

	url, err := u.Parse(urlStr)
	if err != nil {
		log.Fatalln("发生致命错误，请输入正确的链接！！")
	}
	if rf, ok := crawler.RFLimit[url.Hostname()]; ok {
		glc = make(chan interface{}, rf.Concurrent)
		gap = rf.Gap
		log.Printf("该网站对请求频率进行了限制，本程序的并发量限制为%d， 所以耗时会更长一点", rf.Concurrent)
	}

	*crawler.Glc = glc
	*crawler.Gap = gap
}

// doCrawler 控制爬取流程
func doCrawler(urlStr, fileName string) {
	if c, err := crawler.CreateCrawler(urlStr); err == nil {

		log.Println("正在获取章节列表......")

		if chapters, err := c.FetchChapterList(); err == nil {
			log.Println(utils.Green("章节列表已获取"))
			log.Println("正在下载章节内容......")

			// 创建文件
			file, err := os.Create(fileName)
			if err != nil {
				panic(err)
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
					log.Println(utils.Red("Error: " + err.Error() + "\n"))
				}
			}(file)

			// 这里也要限制一下并发量，为什么呢，因为有些章节是分页展示的，如果过这里不限制并发量，所有章节的所有页面都随机地获取
			// 容易出现爬取的页面虽然很多，但爬取的完整章节很少的情况。这时候在前期进度条就会始终显示为0，虽然爬取总时间不变，用户体验感不好。
			glc := make(chan interface{}, 50)
			if rf, ok := crawler.RFLimit[c.GetUrl().Hostname()]; ok {
				glc = make(chan interface{}, rf.Concurrent)
			}

			// 进度条，进度条每次输出时，会把上一行消除掉，所以打日志时每行末尾多加一个\n
			p := mpb.New(mpb.WithWidth(64))
			bar := p.New(int64(len(chapters)),
				// BarFillerBuilder with custom style
				mpb.BarStyle().Lbound("╢").Filler("=").Tip(">").Padding("-").Rbound("╟"),
				mpb.PrependDecorators(
					decor.Name(utils.Green("章节下载中......"), decor.WC{W: len("章节下载中......") + 1, C: decor.DidentRight}),
					decor.Name(utils.Green("进度："), decor.WCSyncSpaceR),
					decor.CountersNoUnit(utils.Green("%d / %d"), decor.WCSyncWidth),
				),
				mpb.AppendDecorators(
					decor.OnComplete(decor.Percentage(decor.WC{W: 5}), "done"),
				),
			)

			// 爬取每一章节的内容
			errChapters := make([]*crawler.Chapter, 0)
			for i := 0; i < len(chapters); i++ {
				go func(idx int) {
					glc <- 1
					defer func() { _ = <-glc }()

					err = retry(func() error {
						return c.FetchChapterContent(&chapters[idx])
					}, 5)
					if err != nil {
						log.Println(utils.Red("Error: " + err.Error() + "\n"))
						errChapters = append(errChapters, &chapters[idx])
					}
					bar.Increment()
				}(i)
			}

			// p 内置waitgroup，也就是等待所有程序爬取完毕
			p.Wait()
			time.Sleep(time.Millisecond * 1000) // 休眠0.1秒，让控制台io同步

			if len(errChapters) > 0 {
				log.Println(utils.Red("由于某些原因，以下章节爬取过程出现错误："))
				for _, ec := range errChapters {
					log.Println(utils.Red(ec.Title))
				}
			}

			log.Println(utils.Green("所有章节爬取完毕......"))
			log.Println("正在把爬取结果写入文件......")
			for _, cha := range chapters {
				err = cha.Save(file)
				if err != nil {
					log.Println(utils.Red("Error: " + err.Error() + "\n"))
				}
			}
			log.Println(utils.Green("程序已运行结束"))
		}

	} else {
		log.Println(utils.Red("Error: " + err.Error() + "\n"))
	}

}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(utils.Yellow("注意，如果程序超过一分钟无响应，请重新执行"))
	var fileName = flag.String("f", "", "保存文件名")
	var urlStr = flag.String("u", "", "url链接")
	flag.Parse()

	initConcurrentLimit(*urlStr)
	doCrawler(*urlStr, *fileName+".txt")

	time.Sleep(time.Second)
}

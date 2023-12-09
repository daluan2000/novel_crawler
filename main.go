package main

import (
	"flag"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
	"log"
	u "net/url"
	"novel_crawler/crawler"
	"novel_crawler/crawler/utils/config_manager"
	"novel_crawler/my_global"
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

func getRFL(hostName string) (chan interface{}, time.Duration) {

	if rfl, ok := my_global.RFLimit[hostName]; ok {
		return make(chan interface{}, rfl.Concurrent), rfl.Gap
	}
	return make(chan interface{}, my_global.DefaultRFL.Concurrent), my_global.DefaultRFL.Gap
}

func initConcurrentLimit(urlStr string) {

	url, err := u.Parse(urlStr)
	if err != nil {
		log.Fatalln("发生致命错误，请输入正确的链接！！")
	}
	crawler.Glc, crawler.Gap = getRFL(url.Hostname())

}

func fetchChapterListWrapper(c crawler.CrawlerInterface) ([]crawler.Chapter, error) {
	chapters := make([]crawler.Chapter, 0)

	err := retry(func() error {
		var err1 error
		chapters, err1 = c.FetchChapterList()
		return err1
	}, 5)

	return chapters, err
}

// doCrawler 控制爬取流程
func doCrawler(urlStr, fileName string) {
	if c, err := crawler.CreateCrawler(urlStr); err == nil {

		log.Printf("本网站爬虫并发量限制为%d", cap(crawler.Glc))

		log.Println("正在获取章节列表......")

		if chapters, err := fetchChapterListWrapper(c); err == nil {
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
			glc, _ := getRFL(c.GetUrl().Hostname())

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
			saveChan := make([]chan interface{}, len(chapters))
			for i := 0; i < len(chapters); i++ {
				saveChan[i] = make(chan interface{}, 1)
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

					saveChan[idx] <- 1
					bar.Increment()
				}(i)
			}

			// 按顺序保存各个章节内容
			for i := 0; i < len(chapters); i++ {
				_ = <-saveChan[i]
				err = chapters[i].Save(file)
				if err != nil {
					log.Println(utils.Red("Error: " + err.Error() + "\n"))
				}
			}

			p.Wait() // 不加这一个，io会不同步

			if len(errChapters) > 0 {
				log.Println(utils.Red("由于某些原因，以下章节爬取过程出现错误："))
				for _, ec := range errChapters {
					log.Println(utils.Red(ec.Title))
				}
			}

			log.Println(utils.Green("所有章节爬取完毕......"))

		}

	} else {
		log.Println(utils.Red("Error: " + err.Error() + "\n"))
	}

}
func readYaml() {
	cm, err := config_manager.CreateConfigManager("yaml", []string{"../", "./"}, my_global.WebInfoFileName)
	if err != nil {
		log.Println(utils.Yellow(my_global.WebInfoFileName + ".yml文件不存在或解析出现了错误，本次爬取使用默认配置。" + err.Error()))
		return
	}

	var uf = cm.GetBool("UseFlag")
	if !uf {
		log.Println(utils.Yellow("UseFlag=false，本次爬取使用默认配置"))
		return
	}

	log.Println(utils.Yellow("本次爬取使用自定义配置文件"))

	var bqg map[string]my_global.BiQuGeInfo
	var nbqg map[string]my_global.NewBiQuGeInfo
	var rfl map[string]my_global.RequestFrequencyLimit

	if cm.Get("BiQuGeInfoByHost") != nil {
		if err = cm.UnmarshalKey("BiQuGeInfoByHost", &bqg); err == nil {
			for k, v := range bqg {
				my_global.BiQuGeInfoByHost[k] = v
			}
		} else {
			log.Println(utils.Red("BiQuGeInfoByHost配置格式错误" + err.Error()))
		}
	}

	if cm.Get("NewBiQuGeInfoByHost") != nil {
		if err = cm.UnmarshalKey("NewBiQuGeInfoByHost", &nbqg); err == nil {
			for k, v := range nbqg {
				my_global.NewBiQuGeInfoByHost[k] = v
			}
		} else {
			log.Println(utils.Red("NewBiQuGeInfoByHost配置格式错误" + err.Error()))
		}
	}

	if cm.Get("RFLimit") != nil {
		if err = cm.UnmarshalKey("RFLimit", &rfl); err == nil {
			for k, v := range rfl {
				my_global.RFLimit[k] = v
			}
		} else {
			log.Println(utils.Red("RFLimit配置格式错误" + err.Error()))
		}
	}
}
func main() {

	my_global.StartTime = time.Now().UnixNano()
	defer func() {
		dur := time.Now().UnixNano() - my_global.StartTime
		log.Printf(utils.Green("程序运行结束，耗时%dmin%ds，共发起%d次请求"),
			dur/time.Minute.Nanoseconds(), dur%time.Minute.Nanoseconds()/time.Second.Nanoseconds(), my_global.RequestCount)
	}()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(utils.Yellow("注意，如果程序超过一分钟无响应，请重新执行"))

	readYaml()

	var fileName = flag.String("f", "", "保存文件名")
	var urlStr = flag.String("u", "", "url链接")
	var saveTitle = flag.Int("st", 1, "保存tittle为1，不保存title为2，不输入该参数默认为1")

	flag.Parse()
	my_global.SaveTitle = *saveTitle == 1

	initConcurrentLimit(*urlStr)
	doCrawler(*urlStr, *fileName+".txt")

	time.Sleep(time.Second)
}

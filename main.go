package main

import (
	"flag"
	"fmt"
	"github.com/redmask-hb/GoSimplePrint/goPrint"
	"log"
	"novel_crawler/crawler"
	"os"
	"sync"
	"time"
)

// 目前适配网站 https://www.52bqg.org/book_128955/

func retry(task func() error, count int) error {
	count--
	if err := task(); err == nil {
		return nil
	} else if count > 0 {
		return retry(task, count-1)
	} else {
		return err
	}

}

// doCrawler 控制爬取流程
func doCrawler(urlStr, fileName string) {
	if c, err := crawler.CreateCrawler(urlStr); err == nil {

		log.Println("正在获取章节列表......")

		if chapters, err := c.FetchChapterList(); err == nil {
			log.Println("章节列表已获取")
			log.Println("正在下载章节内容......")

			// 进度条
			bar := goPrint.NewBar(len(chapters))
			bar.SetNotice("已下载章节：")
			bar.SetGraph(">")

			// 创建文件
			file, err := os.Create(fileName)
			if err != nil {
				panic(err)
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
					log.Println(err.Error())
				}
			}(file)

			// 限制并发量，防止tcp端口耗尽
			limit := make(chan interface{}, 50)
			w := sync.WaitGroup{}
			cnt := 0 // 计数器

			// 爬取每一章节的内容
			for i := 0; i < len(chapters); i++ {
				w.Add(1)
				go func(idx int) {
					limit <- 1
					defer func() { _ = <-limit }()

					err = retry(func() error {
						return c.FetchChapterContent(&chapters[idx])
					}, 5)
					if err != nil {
						log.Println(err.Error())
					}
					w.Done()
					cnt++
					bar.PrintBar(cnt)
				}(i)

			}

			w.Wait()

			time.Sleep(time.Millisecond * 100) // 休眠0.1秒，让控制台io同步
			fmt.Println()
			log.Println("所有章节爬取完毕......")
			log.Println("正在把爬取结果写入文件......")
			for _, cha := range chapters {
				err = cha.Save(file)
				if err != nil {
					log.Println(err.Error())
				}
			}
			log.Println("程序已完成，可以退出")
		}

	} else {
		log.Println(err.Error())
	}

}

func main() {
	log.Println("注意，如果程序超过一分钟无响应，请重新执行")
	var fileName = flag.String("f", "", "保存文件名")
	var urlStr = flag.String("u", "", "url链接")
	flag.Parse()
	doCrawler(*urlStr, *fileName+".txt")
	time.Sleep(time.Hour)
}

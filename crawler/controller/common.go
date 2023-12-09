package controller

import (
	"errors"
	"fmt"
	"log"
	u "net/url"
	"novel_crawler/crawler/chapter/chapter_handler"
	"novel_crawler/crawler/chapter/chapter_interf"
	"novel_crawler/crawler/fetcher_content"
	"novel_crawler/crawler/fetcher_list"
	"novel_crawler/crawler/utils/color_util"
	"novel_crawler/crawler/utils/retry"
	"novel_crawler/crawler/utils/str_util"
	"novel_crawler/global/consts"
	"novel_crawler/global/variable"
	"os"
	"time"
)

type common struct {
}

func (c *common) DoCrawling(url *u.URL, fileName string) {

	st := time.Now().Second()

	// 读取并发限制相关信息
	info := variable.InfoStore.GetInfo(url)
	glc := make(chan interface{}, info.Concurrent)
	gap := info.Gap
	log.Println(color_util.Yellow(fmt.Sprintf("本次爬虫并发量限制为%d", info.Concurrent)))

	// 爬取章节列表，error进行retry
	fl := fetcher_list.Fatory.CreateFetcher(url)
	chapters := make([]chapter_interf.Chapter, 0)
	err := retry.Retry(func() error {
		var err1 error
		chapters, err1 = fl.Fetch(url)
		return err1
	}, consts.RetryCount)
	if err != nil {
		log.Println(color_util.Red(err.Error()))
		return
	}

	// 创建文件
	f, err := os.Create(fileName)
	if err != nil {
		log.Println(color_util.Red(err.Error()))
		return
	}

	p, bar := str_util.ProgressBar(len(chapters))

	// 爬取章节内容
	sc := make([]chan interface{}, len(chapters))
	for i := 0; i < len(sc); i++ {
		sc[i] = make(chan interface{}, 1)
	}
	errChapters := make([]chapter_interf.Chapter, 0)
	for i := 0; i < len(chapters); i++ {
		go func(idx int) {

			// 并发限制
			glc <- 1
			defer func() {
				if gap > 0 {
					time.Sleep(gap)
				}
				_ = <-glc
			}()

			err := errors.New("")
			fc := fetcher_content.Factory.CreateFetcher(chapters[idx].Url)
			ch := chapter_handler.Handler{}
			hasErr := true
			if err = fc.Fetch(&chapters[idx]); err == nil {
				if err = ch.DoBeforeSave(&chapters[idx]); err == nil {
					hasErr = false
				}
			}
			if hasErr {
				chapters[idx].Err = err
				errChapters = append(errChapters, chapters[idx])
				log.Println(color_util.Red(chapters[idx].Title + err.Error()))
			}
			sc[idx] <- 1
		}(i)
		go func(idx int) {
			_ = <-sc[idx]
			ch := chapter_handler.Handler{}
			if err = ch.Save(f, &chapters[idx]); err != nil {
				log.Println(color_util.Red("文件写入错误" + err.Error()))
			}
			bar.Increment()
		}(i)
	}

	p.Wait()

	if len(errChapters) > 0 {
		log.Println(color_util.Red("由于某些原因，下列章节爬取错误"))
		for _, v := range errChapters {
			log.Println(color_util.Red(fmt.Sprintf("%s，错误原因：%s\n", v.Title, v.Err.Error())))
		}
	}

	ed := time.Now().Second()
	min := (ed - st) / 60
	sec := (ed - st) % 60
	log.Println(color_util.Green(fmt.Sprintf("爬取结束，共用时%d分%d秒，共发起%d次http请求", min, sec, variable.RequestCount)))
}

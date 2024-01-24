package controller

import (
	"fmt"
	"log"
	u "net/url"
	"novel_crawler/crawler/chapter/chapter_interf"
	"novel_crawler/crawler/fetcher_content"
	"novel_crawler/crawler/fetcher_list"
	"novel_crawler/crawler/utils/color_util"
	"novel_crawler/crawler/utils/common_util"
	"novel_crawler/global/variable"
	"os"
	"time"
)

type common struct {
}

func (c *common) DoCrawling(url *u.URL, fileName string) {

	st := time.Now().UnixMilli()

	// 读取并发限制相关信息
	info := variable.InfoStore.GetInfo(url)
	glc := make(chan interface{}, info.Concurrent)
	gap := info.Gap
	log.Println(color_util.Yellow(fmt.Sprintf("本次爬虫并发量限制为%d", info.Concurrent)))

	// 爬取章节列表，error进行retry
	log.Println("开始获取章节目录...")
	fl := fetcher_list.Fatory.CreateFetcher(url)
	chapters := make([]chapter_interf.Chapter, 0)
	err := common_util.Retry(func() error {
		var err1 error
		chapters, err1 = fl.Fetch(url)
		return err1
	}, variable.RetryCount)
	if err != nil {
		log.Println(color_util.Red(err.Error()))
		return
	}
	log.Println(color_util.Green("已获取章节目录"))

	// 创建文件
	f, err := os.Create(fileName)
	if err != nil {
		log.Println(color_util.Red(err.Error()))
		return
	}

	p, bar := common_util.ProgressBar(len(chapters))

	// 爬取章节内容
	sc := make([]chan interface{}, len(chapters))
	for i := 0; i < len(sc); i++ {
		sc[i] = make(chan interface{}, 1)
	}
	errChapters := make([]chapter_interf.Chapter, 0)
	for i := 0; i < len(chapters); i++ {
		go func(idx int) {

			// 并发限制
			// requester那里已做并发限制，这里还要再做一次，防止开始阶段分散爬取不同章节的第一页
			glc <- 1
			defer func() {
				if gap > 0 {
					// 这个睡眠是否有必要？
					time.Sleep(gap)
				}
				_ = <-glc
			}()

			fc := fetcher_content.Factory.CreateFetcher(chapters[idx].Url)
			ch := variable.ChapterHandler
			hasErr := true

			// 进行爬取，进行retry操作
			err := common_util.Retry(func() error {
				if err = fc.Fetch(&chapters[idx]); err == nil {
					if err = ch.DoBeforeSave(&chapters[idx]); err == nil {
						hasErr = false
					}
				}
				return err
			}, variable.RetryCount)

			// 如果出错，则记录下来
			if hasErr {
				chapters[idx].Err = err
				errChapters = append(errChapters, chapters[idx])
				log.Println(color_util.Red(chapters[idx].Title + err.Error() + "\n"))
			}
			// 向管道发送信息，表明idx章节可以保存了
			sc[idx] <- 1
		}(i)
	}

	// 保存章节
	for i := 0; i < len(chapters); i++ {
		_ = <-sc[i]
		ch := variable.ChapterHandler
		if err = ch.Save(f, &chapters[i]); err != nil {
			log.Println(color_util.Red("文件写入错误" + err.Error()))
		}
		// 进度条增加
		bar.Increment()
	}
	p.Wait()

	// 打印错误信息
	if len(errChapters) > 0 {
		log.Println(color_util.Red("由于某些原因，下列章节爬取错误"))
		for _, v := range errChapters {
			log.Println(color_util.Red(fmt.Sprintf("%s，错误原因：%s", v.Title, v.Err.Error())))
		}
	}

	// 打印时间信息
	ed := time.Now().UnixMilli()
	st /= 1000
	ed /= 1000
	minute := (ed - st) / 60
	second := (ed - st) % 60
	log.Println(color_util.Green(fmt.Sprintf("爬取结束，共用时%d分%d秒，共发起%d次http请求", minute, second, variable.RequestCount)))
}

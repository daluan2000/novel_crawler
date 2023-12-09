package controller

import (
	"errors"
	"log"
	u "net/url"
	"novel_crawler/crawler/chapter"
	"novel_crawler/crawler/fetcher_content"
	"novel_crawler/crawler/fetcher_list"
	"novel_crawler/crawler/utils/color_util"
	"novel_crawler/crawler/utils/str_util"
	"novel_crawler/global/variable"
	"os"
	"time"
)

type Common struct {
}

func (c *Common) DoCrawling(url *u.URL, fileName string) {

	info := variable.InfoStore.GetInfo(url)
	glc := make(chan interface{}, info.FrequencyLimit.Concurrent)
	gap := info.FrequencyLimit.Gap

	fl := fetcher_list.Fatory.CreateFetcher(url)

	chapters, err := fl.Fetch(url)

	if err != nil {
		log.Println(color_util.Red(err.Error()))
		return
	}

	f, err := os.Create(fileName)
	if err != nil {
		log.Println(color_util.Red(err.Error()))
		return
	}

	p, bar := str_util.ProgressBar(len(chapters))

	errChapters := make([]chapter.Chapter, 0)
	for i := 0; i < len(chapters); i++ {
		go func(idx int) {
			glc <- 1
			defer func() {
				if gap > 0 {
					time.Sleep(gap)
				}
				_ = <-glc
			}()

			err := errors.New("")
			fc := fetcher_content.Factory.CreateFetcher(chapters[idx].Url)
			hasErr := true
			if err = fc.Fetch(&chapters[idx]); err == nil {
				if err = chapters[idx].DoBeforeSave(); err == nil {
					if err = chapters[idx].Save(f); err == nil {
						hasErr = false
					}
				}
			}
			if hasErr {
				chapters[idx].Err = err
				errChapters = append(errChapters, chapters[idx])
			}
			bar.Increment()
		}(i)
	}

	p.Wait()

	log.Println("爬取结束")

}

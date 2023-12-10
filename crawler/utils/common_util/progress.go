package common_util

import (
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
	"novel_crawler/crawler/utils/color_util"
)

func ProgressBar(count int) (*mpb.Progress, *mpb.Bar) {
	// 进度条，进度条每次输出时，会把上一行消除掉，所以打日志时每行末尾多加一个\n
	p := mpb.New(mpb.WithWidth(64))
	bar := p.New(int64(count),
		// BarFillerBuilder with custom style
		mpb.BarStyle().Lbound("╢").Filler("=").Tip(">").Padding("-").Rbound("╟"),
		mpb.PrependDecorators(
			decor.Name(color_util.Green("章节下载中......"), decor.WC{W: len("章节下载中......") + 1, C: decor.DidentRight}),
			decor.Name(color_util.Green("进度："), decor.WCSyncSpaceR),
			decor.CountersNoUnit(color_util.Green("%d / %d"), decor.WCSyncWidth),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.Percentage(decor.WC{W: 5}), "done"),
		),
	)
	return p, bar
}

package info

import (
	"novel_crawler/crawler/info/info_interf"
	"time"
)

/*
网站的一切信息存在此文件的数据结构里
包括选择器、filter、fetcher等等
*/
var infoMap = map[string]info_interf.Info{
	// 笔趣阁
	"www.2biqu.com": {
		ASelector:       ".section-list > li > a",
		ContentSelector: "#content",
		FrequencyLimit: info_interf.FrequencyLimit{
			Concurrent: 4,
			Gap:        time.Millisecond * 250,
		},
	},

	// 笔趣阁
	"www.bige3.cc": {
		ASelector:       "dl dd > a:not(a[rel='nofollow'])",
		ContentSelector: "#chaptercontent",
		RemoveSelector:  []string{"p"},
	},

	// 笔趣阁 该网站搜索时会进行人机检测，防止人机验证加载不出来，最好使用chrome浏览器，
	"www.52bqg.org": {
		ASelector:       "dd > a",
		ContentSelector: "#content",
	},

	"www.ujxsw.net": {
		ASelector:       "#readerlist > ul > li > a",
		ContentSelector: ".read-content > p",
	},
	// 天涯读书，有一些出版读物
	"www.tianyabook.com": {
		ASelector:       ".panel-body > dd > a",
		ContentSelector: "#htmlContent",
		RemoveSelector:  []string{"p"},
	},
	// 同人小说网，二次元书籍
	"www.trxs.cc": {
		ASelector:       ".book_list > ul > li > a",
		ContentSelector: ".read_chapterDetail",
		RemoveSelector:  []string{"img"},
	},
	// 科幻小说网
	"www.00txt.com": {
		ASelector:       ".list-group > li.vv-book > a",
		ContentSelector: "#content",
		RemoveSelector:  []string{"div"},
	},
	// 好笔阁
	"www.1688by.com": {
		ASelector:       "#list > dl > a",
		ContentSelector: "#booktxt",
		RemoveSelector:  []string{"div", "p[style*='color']"},
	},

	// 好笔阁
	"www.bixiashenghua.com": {
		ASelector:       "#list > dl > dd > a",
		ContentSelector: "#content",
	},

	// 千叶阁 sb网站限制频率
	"www.qianyege.com": {
		ASelector:       "#list > dl > dd > a",
		ContentSelector: "#content",
		RemoveSelector:  []string{"div"},
		FrequencyLimit: info_interf.FrequencyLimit{
			Concurrent: 4,
			Gap:        time.Millisecond * 250,
		},
	},
	// 笔趣阁
	"www.biquinfo.com": {
		ASelector:       "#section-list > li > a",
		ContentSelector: "#content",
	},

	"www.wbsz.org": {
		ASelector:       ".chapter > ul > li > a",
		ContentSelector: ".readerCon",
		RemoveSelector:  []string{"script"},
		FrequencyLimit: info_interf.FrequencyLimit{
			Concurrent: 4,
			Gap:        time.Millisecond * 250,
		},
	},
	"www.beqege.com": {
		ASelector:       "#list > dl > dd > a",
		ContentSelector: "#content",
		FrequencyLimit: info_interf.FrequencyLimit{
			Concurrent: 4,
			Gap:        time.Millisecond * 250,
		},
	},
	"www.xsbiquge.la": {
		ASelector:       ".listmain > dl > dd > a",
		ContentSelector: "#content",
		StrReplace: map[string]string{
			"<p class=\"content_detail\">": "",
		},
	},
	/*------------------------------------分割线-------------------------------------------------------------------------------------*/
	/*------------------------------------分割线-------------------------------------------------------------------------------------*/
	/*------------------------------------分割线-------------------------------------------------------------------------------------*/

	"www.xbiqugeo.com": {

		NextChapterList: info_interf.NextChapterList{
			MultiPageChapterList:    true,
			ChapterListNextSelector: ".listpage > .right > a",
			ChapterListNextStr:      "下一页",
		},

		NextContent: info_interf.NextContent{
			MultiPageContent:    true,
			ContentNextStr:      "下一页",
			ContentNextSelector: "#next_url",
		},
		ASelector:       ".section-box:nth-child(4) > ul > li > a",
		ContentSelector: "#content",
		RemoveSelector:  []string{"a", "div"},
	},

	"www.zrfsxs.com": {

		NextChapterList: info_interf.NextChapterList{
			MultiPageChapterList:    true,
			ChapterListNextSelector: "#pages > a.gr",
			ChapterListNextStr:      "下一页",
		},

		NextContent: info_interf.NextContent{
			MultiPageContent:    true,
			ContentNextStr:      "下一页",
			ContentNextSelector: ".prenext > span:nth-child(3) > a",
		},

		ASelector:       "#list > ul > li > a",
		ContentSelector: ".con",
	},

	"youyouxs.com": {
		NextChapterList: info_interf.NextChapterList{
			MultiPageChapterList:    true,
			ChapterListNextSelector: ".index-container-btn:last-child",
			ChapterListNextStr:      "下一页",
		},

		NextContent: info_interf.NextContent{
			MultiPageContent:    true,
			ContentNextStr:      "下一页",
			ContentNextSelector: ".bottem1 > a:last-child",
		},

		ASelector:       "a[rel='chapter']",
		ContentSelector: "#booktxt",

		RemoveSelector: []string{"div"},
	},

	"www.biqge.org": {

		NextChapterList: info_interf.NextChapterList{
			MultiPageChapterList:    true,
			ChapterListNextSelector: "a.index-container-btn:last-child",
			ChapterListNextStr:      "下一页",
		},

		NextContent: info_interf.NextContent{
			MultiPageContent:    true,
			ContentNextStr:      "下一页",
			ContentNextSelector: "#next_url",
		},

		ASelector:       ".section-box:nth-child(4) li > a",
		ContentSelector: "#content",
	},

	// 和上面那个一模一样
	"www.ddxs.vip": {
		NextChapterList: info_interf.NextChapterList{
			MultiPageChapterList:    true,
			ChapterListNextSelector: "a.index-container-btn:last-child",
			ChapterListNextStr:      "下一页",
		},

		NextContent: info_interf.NextContent{
			MultiPageContent:    true,
			ContentNextStr:      "下一页",
			ContentNextSelector: "#next_url",
		},

		ASelector:       ".section-box:nth-child(4) li > a",
		ContentSelector: "#content",
	},
}

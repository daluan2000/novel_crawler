package info

import (
	"novel_crawler/crawler/info/info_interf"
	"time"
)

func AddInfo(im map[string]info_interf.Info) {
	for k, v := range im {
		infoMap[k] = v
	}
}

var infoMap = map[string]info_interf.Info{
	// 笔趣阁
	"www.2biqu.com": {
		ASelector:       ".section-list > li > a",
		ContentSelector: "#content",
		StrReplace: map[string]string{
			"\u807d": " ",
			"<br>":   "",
			"<br/>":  "",
		},
		FrequencyLimit: info_interf.FrequencyLimit{
			Concurrent: 4,
			Gap:        time.Millisecond * 250,
		},
	},

	// 笔趣阁
	"www.bige3.cc": {
		ASelector:       "dl dd > a:not(a[rel='nofollow'])",
		ContentSelector: "#chaptercontent",
		StrReplace: map[string]string{
			"<br/><br/>": "\n",
			"<br><br>":   "\n",
		},
		RemoveSelector: []string{"p"},
	},

	// 笔趣阁 该网站搜索时会进行人机检测，防止人机验证加载不出来，最好使用chrome浏览器，
	"www.52bqg.org": {
		ASelector:       "dd > a",
		ContentSelector: "#content",
		StrReplace: map[string]string{
			"\u807d聽聽聽":  "    ",
			"<br/><br/>": "\n",
			"<br><br>":   "\n",
		},
	},

	"www.ujxsw.net": {
		ASelector:       "#readerlist > ul > li > a",
		ContentSelector: ".read-content > p",
		StrReplace: map[string]string{
			"<br>":         "",
			"<br/>\n<br/>": "",
		},
	},
	// 天涯读书，有一些出版读物
	"www.tianyabook.com": {
		ASelector:       ".panel-body > dd > a",
		ContentSelector: "#htmlContent",
		StrReplace: map[string]string{
			"聽":     "",
			"<br>":  "",
			"<br/>": "",
		},
		RemoveSelector: []string{"p"},
	},
	// 同人小说网，二次元书籍
	"www.trxs.cc": {
		ASelector:       ".book_list > ul > li > a",
		ContentSelector: ".read_chapterDetail",
		StrReplace: map[string]string{
			"<p>":        "\n",
			"</p>":       "",
			"<br/><br/>": "\n",
		},
		RemoveSelector: []string{"img"},
	},
	// 科幻小说网
	"www.00txt.com": {
		ASelector:       ".list-group > li.vv-book > a",
		ContentSelector: "#content",
		StrReplace: map[string]string{
			"<p>":   "",
			"</p>":  "",
			"<br/>": "",
			"<br>":  "",
		},
		RemoveSelector: []string{"div"},
	},
	// 好笔阁
	"www.1688by.com": {
		ASelector:       "#list > dl > a",
		ContentSelector: "#booktxt",
		StrReplace: map[string]string{
			"<p>":   "\n",
			"</p>":  "",
			"<br/>": "",
			"<br>":  "",
		},
		RemoveSelector: []string{"div", "p[style*='color']"},
	},

	// 好笔阁
	"www.bixiashenghua.com": {
		ASelector:       "#list > dl > dd > a",
		ContentSelector: "#content",
		StrReplace: map[string]string{
			"<p>":   "\n",
			"</p>":  "",
			"<br/>": "",
			"<br>":  "",
		},
		RemoveSelector: []string{},
	},

	// 千叶阁 sb网站限制频率
	"www.qianyege.com": {
		ASelector:       "#list > dl > dd > a",
		ContentSelector: "#content",
		StrReplace: map[string]string{
			"<p>":   "",
			"</p>":  "",
			"<br/>": "",
			"<br>":  "",
			"聽":     " ",
		},
		RemoveSelector: []string{"div"},
		FrequencyLimit: info_interf.FrequencyLimit{
			Concurrent: 4,
			Gap:        time.Millisecond * 250,
		},
	},
	// 笔趣阁
	"www.biquinfo.com": {
		ASelector:       "#section-list > li > a",
		ContentSelector: "#content",
		StrReplace: map[string]string{
			"<br/>": "",
			"<br>":  "",
			"聽":     " ",
		},
		RemoveSelector: []string{},
	},

	"www.wbsz.org": {
		ASelector:       ".chapter > ul > li > a",
		ContentSelector: ".readerCon",
		StrReplace: map[string]string{
			"<br/>":   "",
			"<br>":    "",
			"<p>":     "",
			"</p>":    "",
			"<span>":  "",
			"</span>": "",
			"聽":       " ",
		},
		RemoveSelector: []string{"script"},
		FrequencyLimit: info_interf.FrequencyLimit{
			Concurrent: 4,
			Gap:        time.Millisecond * 250,
		},
	},
	"www.beqege.com": {
		ASelector:       "#list > dl > dd > a",
		ContentSelector: "#content",
		StrReplace: map[string]string{
			"<br/>":   "",
			"<br>":    "",
			"<p>":     "",
			"</p>":    "",
			"<span>":  "",
			"</span>": "",
			"聽":       " ",
		},
		RemoveSelector: []string{},
		FrequencyLimit: info_interf.FrequencyLimit{
			Concurrent: 4,
			Gap:        time.Millisecond * 250,
		},
	},

	/*------------------------------------分割线-------------------------------------------------------------------------------------*/
	/*------------------------------------分割线-------------------------------------------------------------------------------------*/
	/*------------------------------------分割线-------------------------------------------------------------------------------------*/

	"www.xbiqugeo.com": {

		NextChapterList: info_interf.NextChapterList{
			HasNextChapterList:      true,
			ChapterListNextSelector: ".listpage > .right > a",
			ChapterListNextStr:      "下一页",
		},

		NextContent: info_interf.NextContent{
			HasNextContent:      true,
			ContentNextStr:      "下一页",
			ContentNextSelector: "#next_url",
		},
		ASelector:       ".section-box:nth-child(4) > ul > li > a",
		ContentSelector: "#content",

		StrReplace: map[string]string{
			"<p>":  "\n    ",
			"</p>": "",
		},
		RemoveSelector: []string{"a", "div"},
	},

	"www.zrfsxs.com": {

		NextChapterList: info_interf.NextChapterList{
			HasNextChapterList:      true,
			ChapterListNextSelector: "#pages > a.gr",
			ChapterListNextStr:      "下一页",
		},

		NextContent: info_interf.NextContent{
			HasNextContent:      true,
			ContentNextStr:      "下一页",
			ContentNextSelector: ".prenext > span:nth-child(3) > a",
		},

		ASelector:       "#list > ul > li > a",
		ContentSelector: ".con",

		StrReplace: map[string]string{
			"<p>":  "\n    ",
			"</p>": "",
		},
	},
}

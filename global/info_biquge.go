package global

type BiQuGeInfo struct {
	ASelector       string
	ContentSelector string

	// 关于字符串替换，有几点要注意：
	// 1. 代码文件的编码格式为utf-8，小说文件的编码格式也应该为utf-8，这样才能保证替换的结果是正确的
	// 2. UTF-8是Unicode的一种实现方式，某些非utf-8编码的字符，在进行编码转换后可能出现乱码现象，这时直接百度搜索对应的unicode编码，然后替换即可
	//    比如对于GBK编码下的&nbsp空格字符，在转换为utf8后它的显示为 聽 \u807d，下面一行做了替换
	StrReplace map[string]string
	// 字符串中删除一些标签
	RemoveSelector []string
}

// http://www.zwduxs.com/102_102828/ 此网站下载会有残缺，不知为何

var BiQuGeInfoByHost = map[string]BiQuGeInfo{
	// 笔趣阁
	"www.2biqu.com": {
		ASelector:       ".section-list > li > a",
		ContentSelector: "#content",
		StrReplace: map[string]string{
			"\u807d": " ",
			"<br>":   "",
			"<br/>":  "",
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
}
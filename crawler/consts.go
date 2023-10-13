package crawler

var hostASelector = map[string]string{
	"www.bige3.cc":       "dl dd > a:not(a[rel='nofollow'])",
	"www.52bqg.org":      "dd > a",
	"www.beqege.com":     "dl > dd > a",
	"www.biquge7.xyz":    ".list > ul > li > a",
	"www.biqigewx.com/":  "#list > dl > dd > a",
	"www.tianyabook.com": ".panel-chapterlist > dd > a",
}

var hostContentSelector = map[string]string{
	"www.bige3.cc":       "#chaptercontent",
	"www.52bqg.org":      "#content",
	"www.beqege.com":     "#content",
	"www.biquge7.xyz":    ".text",
	"www.biqigewx.com/":  "#content",
	"www.tianyabook.com": "#htmlContent",
}

var hostReplace = map[string][]string{
	"www.bige3.cc": {
		"<br/><br/>", "\n",
		"<br><br>", "\n",
	},

	// UTF-8是Unicode的一种实现方式，对于中文乱码，直接百度搜索对应的unicode编码，然后替换即可
	// 对于&nbsp空格，他的utf-8是聽，下面一行做了替换
	"www.52bqg.org": {
		"\u807d聽聽聽", "    ",
		"<br/><br/>", "\n",
		"<br><br>", "\n",
	},

	"www.beqege.com": {
		"<p>", "",
		"</p>", "",
	},

	"www.biquge7.xyz": {
		"\u807d聽聽聽", "    ",
		"<br/><br/>", "\n",
		"<br><br>", "\n",
	},

	"www.biqigewx.com/": {
		"\u807d聽聽聽", "    ",
		"<br/>", "",
		"<br>", "",
	},

	"www.tianyabook.com": {
		"<p class=\"booktag\">", "",
		"</p>", "",
		"</b>", "",
		"<b>", "",
		"\u807d聽聽聽", "    ",
		"<br/>", "",
		"<br>", "",
	},
}

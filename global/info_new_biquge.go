package global

type NewBiQuGeInfo struct {
	// 章节列表下一页a标签的选择器，章节内容下一页a标签的选择器
	ChapterListNextSelector string
	ContentNextSelector     string

	// 章节列表下一页a标签应包含的文本，章节内容下一页a标签应包含的文本
	ChapterListNextStr string
	ContentNextStr     string

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

var NewBiQuGeInfoByHost = map[string]NewBiQuGeInfo{
	"www.xbiqugeo.com": {
		ChapterListNextSelector: ".listpage > .right > a",
		ContentNextSelector:     "#next_url",

		ASelector:       ".section-box:nth-child(4) > ul > li > a",
		ContentSelector: "#content",

		ChapterListNextStr: "下一页",
		ContentNextStr:     "下一页",

		StrReplace: map[string]string{
			"<p>":  "\n    ",
			"</p>": "",
		},
		RemoveSelector: []string{"a", "div"},
	},

	"www.zrfsxs.com": {
		ChapterListNextSelector: "#pages > a.gr",
		ContentNextSelector:     ".prenext > span:nth-child(3) > a",

		ASelector:       "#list > ul > li > a",
		ContentSelector: ".con",

		ChapterListNextStr: "下一页",
		ContentNextStr:     "下一页",

		StrReplace: map[string]string{
			"<p>":  "\n    ",
			"</p>": "",
		},
	},

	"youyouxs.com": {
		ChapterListNextSelector: ".index-container-btn:last-child",
		ContentNextSelector:     ".bottem1 > a:last-child",

		ASelector:       "a[rel='chapter']",
		ContentSelector: "#booktxt",

		ChapterListNextStr: "下一页",
		ContentNextStr:     "下一页",

		StrReplace: map[string]string{
			"<p>":  "\n    ",
			"</p>": "",
		},

		RemoveSelector: []string{"div"},
	},
}

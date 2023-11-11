package global

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

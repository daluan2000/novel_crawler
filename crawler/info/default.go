package info

import (
	"novel_crawler/crawler/info/info_interf"
	"novel_crawler/crawler/utils/str_util"
	"time"
)

var baseStrReplace = map[string]string{
	"聽":      "",
	" ":      "",
	"\u0010": "",
}
var baseRegReplace = map[string]string{
	`<br[\s/]*?>`:                 "\n",
	str_util.TagRegexp("div")[0]:  "\n",
	str_util.TagRegexp("div")[1]:  "\n",
	str_util.TagRegexp("span")[0]: "\n",
	str_util.TagRegexp("span")[1]: "\n",
	str_util.TagRegexp("p")[0]:    "\n",
	str_util.TagRegexp("p")[1]:    "\n",
}

var defaultFL = info_interf.FrequencyLimit{
	Concurrent: 50,
	Gap:        time.Millisecond * 0,
}

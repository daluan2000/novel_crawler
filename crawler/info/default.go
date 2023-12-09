package info

import (
	"novel_crawler/crawler/info/info_interf"
	"time"
)

var baseReplace = map[string]string{
	"聽":      "",
	"</br>":   "\n",
	"<br>":    "\n",
	"<p>":     "\n",
	"</p>":    "\n",
	"<span>":  "\n",
	"</span>": "\n",
	"<div>":   "\n",
	"</div>":  "\n",
	" ":       "",
}
var defaultRFL = info_interf.FrequencyLimit{
	Concurrent: 50,
	Gap:        time.Millisecond * 0,
}

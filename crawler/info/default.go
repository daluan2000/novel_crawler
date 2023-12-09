package info

import (
	"novel_crawler/crawler/info/info_interf"
	"time"
)

var defaultRFL = info_interf.FrequencyLimit{
	Concurrent: 50,
	Gap:        time.Millisecond * 0,
}

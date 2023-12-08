package getter_info

import (
	"novel_crawler/crawler/getter_info/getter_info_interf"
	"time"
)

var defaultRFL = getter_info_interf.FrequencyLimit{
	Concurrent: 50,
	Gap:        time.Millisecond * 0,
}

package my_global

import "time"

type RequestFrequencyLimit struct {
	// 并发量限制
	Concurrent int
	// 每次请求后线程的休眠时间
	Gap time.Duration
}

var DefaultRFL = RequestFrequencyLimit{
	Concurrent: 50,
	Gap:        time.Millisecond * 0,
}
var RFLimit = map[string]RequestFrequencyLimit{
	"youyouxs.com": {
		Concurrent: 1,
		Gap:        time.Millisecond * 0,
	},
	"www.qianyege.com": {
		Concurrent: 4,
		Gap:        time.Millisecond * 250,
	},
	"www.2biqu.com": {
		Concurrent: 4,
		Gap:        time.Millisecond * 250,
	},
	"www.biqge.org": {
		Concurrent: 4,
		Gap:        time.Millisecond * 250,
	},
}

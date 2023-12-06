package getter_next

import (
	u "net/url"
	"novel_crawler/crawler/getter_next/getter_next_interf"
)

func CreateContentNextGetter(url *u.URL) getter_next_interf.Getter {
	return &CommonContent{}
}

func CreateChapterListNextGetter(url *u.URL) getter_next_interf.Getter {
	return &CommonChapterList{}
}

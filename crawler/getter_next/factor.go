package getter_next

import u "net/url"

func CreateContentNextGetter(url *u.URL) Getter {
	return &CommonContent{}
}

func CreateChapterListNextGetter(url *u.URL) Getter {
	return &CommonChapterList{}
}

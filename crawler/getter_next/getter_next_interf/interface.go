package getter_next_interf

import (
	"github.com/PuerkitoBio/goquery"
	u "net/url"
)

type Getter interface {
	NextUrl(dom *goquery.Document) (*u.URL, error)
}
type Factory interface {
	CreateContentNextGetter(url *u.URL) Getter
	CreateChapterListNextGetter(url *u.URL) Getter
}

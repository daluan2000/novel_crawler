package getter_next

import (
	"github.com/PuerkitoBio/goquery"
	u "net/url"
)

type Getter interface {
	NextUrl(dom *goquery.Document) (*u.URL, error)
}

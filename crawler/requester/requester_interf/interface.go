package requester_interf

import (
	"github.com/PuerkitoBio/goquery"
	u "net/url"
)

type Requester interface {
	CreateGoQuery(url *u.URL) (*goquery.Document, error)
}
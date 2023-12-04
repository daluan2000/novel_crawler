package http_query

import (
	"github.com/PuerkitoBio/goquery"
	u "net/url"
)

type Query interface {
	CreateGoQuery(url *u.URL) (*goquery.Document, error)
}

package chapter_handler

import (
	u "net/url"
	"novel_crawler/crawler/chapter/chapter_interf"
)

var Factory factory

type factory struct {
}

func (f *factory) CreateHandler(url *u.URL) chapter_interf.Handler {
	return &Handler{}
}

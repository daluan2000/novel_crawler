package controller

import (
	u "net/url"
	"novel_crawler/crawler/controller/controller_interf"
)

var Factory controller_interf.Factory = &factory{}

type factory struct {
}

func (f *factory) CreateController(url *u.URL) controller_interf.Controller {
	return &common{}
}

package controller_interf

import u "net/url"

type Controller interface {
	DoCrawling(url *u.URL, fileName string)
}

type Factory interface {
	CreateController(url *u.URL) Controller
}

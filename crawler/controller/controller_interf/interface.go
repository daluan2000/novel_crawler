package controller_interf

import u "net/url"

type Controller interface {
}

type Factory interface {
	CreateController(url *u.URL) Controller
}

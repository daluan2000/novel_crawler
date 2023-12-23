package chapter_interf

import (
	u "net/url"
	"os"
)

type Handler interface {
	Save(f *os.File, c *Chapter) error
	DoBeforeSave(c *Chapter) error
}

type HandlerFactory interface {
	CreateHandler(url *u.URL) Handler
}

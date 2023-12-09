package bootstrap

import (
	"log"
	"novel_crawler/crawler/info"
	"novel_crawler/global/variable"
)

func init() {
	log.SetFlags(log.LstdFlags)
	variable.InfoStore = info.CreateStore()
}

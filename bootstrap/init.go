package bootstrap

import (
	"novel_crawler/crawler/info"
	"novel_crawler/global/variable"
)

func init() {
	if err := info.ReadYaml(); err != nil {

	}
	variable.InfoStore = info.CreateStore()
}

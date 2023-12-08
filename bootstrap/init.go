package bootstrap

import (
	"novel_crawler/crawler/getter_info"
	"novel_crawler/global/variable"
)

func init() {
	if err := getter_info.ReadYaml(); err != nil {

	}
	variable.GetterInfo = getter_info.CreateGetter()
}

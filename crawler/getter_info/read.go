package getter_info

import (
	"novel_crawler/config_manager"
	"novel_crawler/crawler/getter_info/getter_info_interf"
	"novel_crawler/global/consts"
)

func addInfo(im map[string]getter_info_interf.Info) {
	for k, v := range im {
		infoMap[k] = v
	}
}

// ReadYaml 从配置文件里读取配置信息
func ReadYaml() error {
	cm, err := config_manager.CreateConfigManager("yaml", []string{"../", "./"}, consts.InfoFileName)
	if err != nil {
		return err
	}

	infos := make(map[string]getter_info_interf.Info)

	if cm.Get("Info") != nil {
		if err = cm.UnmarshalKey("Info", &infos); err == nil {
			addInfo(infos)
			return nil
		}
	}
	return err
}

package info

import (
	"novel_crawler/crawler/info/info_interf"
	"novel_crawler/crawler/utils/config_manager"
)

func addInfo(im map[string]info_interf.Info) {
	for k, v := range im {
		infoMap[k] = v
	}
}

// ReadYaml 从配置文件里读取配置信息
func ReadYaml(fileName string) error {
	cm, err := config_manager.CreateConfigManager("yaml", []string{"../", "./"}, fileName)
	if err != nil {
		return err
	}

	infos := make(map[string]info_interf.Info)

	if cm.Get("Info") != nil {
		if err = cm.UnmarshalKey("Info", &infos); err == nil {
			addInfo(infos)
			return nil
		}
	}
	return err
}

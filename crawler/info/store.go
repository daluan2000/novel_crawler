package info

import (
	u "net/url"
	"novel_crawler/crawler/info/info_interf"
	"novel_crawler/crawler/utils/config_manager"
)

type store struct {
}

// ReadYaml 读取yaml配置文件，注意一定要在 FillInfoDefault 之前调用
func (s *store) ReadYaml(fileName string) error {
	cm, err := config_manager.CreateConfigManager("yaml", []string{"../", "./"}, fileName)
	if err != nil {
		return err
	}

	infos := make(map[string]info_interf.Info)

	if cm.Get("Info") != nil {
		if err = cm.UnmarshalKey("Info", &infos); err == nil {
			// 把读取到的文件添加到infomap里
			for k, v := range infos {
				infoMap[k] = v
			}
			return nil
		}
	}
	return err
}

// FillInfoDefault 填充default和base值
func (s *store) FillInfoDefault() {
	filledInfoMap := make(map[string]info_interf.Info)
	for host, info := range infoMap {
		if info.FrequencyLimit.Concurrent == 0 {
			info.FrequencyLimit = defaultFL
		}
		if info.RemoveSelector == nil {
			info.RemoveSelector = make([]string, 0)
		}
		if info.StrReplace == nil {
			info.StrReplace = make(map[string]string)
		}
		if info.RegReplace == nil {
			info.RegReplace = make(map[string]string)
		}

		for k, v := range baseRegReplace {
			info.RegReplace[k] = v
		}
		for k, v := range baseStrReplace {
			info.StrReplace[k] = v
		}

		// 因为遍历过程中不能修改map，所以要另存一个map里
		filledInfoMap[host] = info
	}
	infoMap = filledInfoMap
}

func (s *store) GetInfoByHost(host string) info_interf.Info {
	res := infoMap[host]
	if v, ok := sameWith[host]; ok {
		res = infoMap[v.Host]
		if v.FrequencyLimit.Concurrent != 0 {
			res.FrequencyLimit = v.FrequencyLimit
		}
	}
	return res
}

func (s *store) GetInfo(url *u.URL) info_interf.Info {
	return s.GetInfoByHost(url.Hostname())
}

func (s *store) Exist(url *u.URL) bool {
	_, ok := infoMap[url.Hostname()]
	return ok
}

// CreateStore 这个比较简单，就不建工厂接口了
func CreateStore() info_interf.InfoStore {
	return &store{}
}

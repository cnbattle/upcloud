package config

import (
	"encoding/json"
	"github.com/cnbattle/upcloud/core/utils"
	"io/ioutil"
)

// Conf 配置
var Conf []ProjectConfig

// ProjectConfig ProjectConfig单配置 struct
type ProjectConfig struct {
	ProjectName string            `json:"project_name"`
	Platform    string            `json:"platform"`
	Path        string            `json:"path"`
	Args        map[string]string `json:"args"`
}

func init() {
	configFile := utils.GetConfig()
	if utils.IsExist(configFile) {
		// 存在读取
		//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			panic(err)
		}
		//读取的数据为json格式，需要进行解码
		err = json.Unmarshal(data, &Conf)
		if err != nil {
			panic(err)
		}
	} else {
		b, err := json.Marshal(Conf)
		if err != nil {
			panic(err)
		}
		//生成json文件
		err = ioutil.WriteFile(configFile, b, 0777)
		if err != nil {
			panic(err)
		}
	}
}

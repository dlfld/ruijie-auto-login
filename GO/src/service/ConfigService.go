package service

import (
	"GO/src/pojo"
	"GO/src/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func writeYaml(src string, config pojo.ConfigData) {
	data, _ := yaml.Marshal(config) // 第二个表示每行的前缀，这里不用，第三个是缩进符号，这里用tab
	ioutil.WriteFile(src, data, 0777)
}

// DetectConfig 检测是否有配置文件，如果有的话就返回true，如果没有的话就创建配置文件并返回false
func DetectConfig() bool {
	src := "configuration.yaml"
	exists, _ := utils.FileExists(src)
	if exists {
		return true
	}

	print("未检测到配置文件，正在生成配置文件。。。。\n")

	configData := &pojo.ConfigData{
		UserId:       "你的学号",
		Password:     "对应的密码",
		Server:       "# 1 移动 2联通 3电信 4校园网",
		TimeInterval: 600,
		LogPath:      "ruijie",
		LogClearDay:  1,
	}
	data, _ := yaml.Marshal(configData)
	ioutil.WriteFile(src, data, 0777)
	return false
}

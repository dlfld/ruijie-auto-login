package service

import (
	"GO/src/pojo"
	"GO/src/utils"
)

// DetectConfig 检测是否有配置文件，如果有的话就返回true，如果没有的话就创建配置文件并返回false
func DetectConfig() bool {
	exists, _ := utils.FileExists("configuration.yaml")
	if exists {
		return true
	}

	print("未检测到配置文件，正在生成配置文件。。。。")
	configData := &pojo.ConfigData{
		UserId:       "你的学号",
		Password:     "对应的密码",
		Server:       "# 1 移动 2联通 3电信 4校园网",
		TimeInterval: 3,
		LogPath:      "ruijie",
		LogSaveDay:   1,
		LogClearDay:  1,
	}
	//configData
}

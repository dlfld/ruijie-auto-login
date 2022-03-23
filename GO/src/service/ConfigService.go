package service

import (
	"GO/src/pojo"
	"GO/src/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// DetectConfig 检测是否有配置文件，如果有的话就返回true，如果没有的话就创建配置文件并返回false
func DetectConfig() bool {
	abs_app, err := os.Executable() // get application location
	if err != nil {
		log.Fatalln(err)
	}
	abs_wd, err := filepath.EvalSymlinks(filepath.Dir(abs_app)) // get floder has the executabel application
	if err != nil {
		log.Fatalln(err)
	}
	src := abs_wd + "/configuration.yaml"
	exists, err := utils.FileExists(src)
	if err != nil {
		log.Fatalln(err)
	}
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
	data, err := yaml.Marshal(configData)
	if err != nil {
		log.Fatalln(err)
	}
	ioutil.WriteFile(src, data, 0777)
	return false
}

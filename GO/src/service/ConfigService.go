package service

import (
	"GO/src/pojo"
	"GO/src/utils"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
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
		UserId:       "",
		Password:     "",
		Server:       "",
		TimeInterval: 60,
		LogPath:      "ral.log",
		LogClearDay:  1,
	}
	data, err := yaml.Marshal(configData)
	if err != nil {
		log.Fatalln(err)
	}
	ioutil.WriteFile(src, data, 0777)
	return false
}

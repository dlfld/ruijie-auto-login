package utils

import (
	"GO/src/pojo"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

// ReadConfig 读取配置文件为结构体
func ReadConfig() pojo.ConfigData {
	var ConfigPojo pojo.ConfigData
	wd, _ := os.Getwd()
	dc := path.Join(wd, "configuration.yaml")
	config, err := ioutil.ReadFile(dc)
	if err != nil {
		fmt.Print(err)
	}
	err1 := yaml.Unmarshal(config, &ConfigPojo)
	if err1 != nil {
		fmt.Println("error")
	}
	return ConfigPojo
}

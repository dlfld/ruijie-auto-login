package utils

import (
	"GO/pojo"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// ReadConfig 读取配置文件为结构体
func ReadConfig() pojo.ConfigData {
	var ConfigPojo pojo.ConfigData
	config, err := ioutil.ReadFile("/Users/dailinfeng/Desktop/小项目/锐捷网页认证登录/GO/configuration.yaml")
	if err != nil {
		fmt.Print(err)
	}
	err1 := yaml.Unmarshal(config, &ConfigPojo)
	if err1 != nil {
		fmt.Println("error")
	}
	return ConfigPojo
}

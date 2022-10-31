package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"auto_login/pojo"

	"gopkg.in/yaml.v2"
)

// ReadConfig 读取配置文件为结构体
func ReadConfig() pojo.ConfigData {
	var ConfigPojo pojo.ConfigData

	absApp, err := os.Executable() // get application location
	if err != nil {
		log.Fatalln(err)
	}
	absWd, _ := filepath.EvalSymlinks(filepath.Dir(absApp)) // get floder has the executabel application

	dc := path.Join(absWd, "configuration.yaml")

	config, err := ioutil.ReadFile(dc)
	if err != nil {
		fmt.Println(err)
	}
	err1 := yaml.Unmarshal(config, &ConfigPojo)
	if err1 != nil {
		fmt.Println(err1)
	}

	// empty LogPath
	if ConfigPojo.LogPath == "" {
		ConfigPojo.LogPath = "ruijie.log"
	}
	ConfigPojo.LogPath = path.Join(absWd, ConfigPojo.LogPath)

	// empty Server
	if ConfigPojo.Server == "" {
		ConfigPojo.Server = "edu"
	}

	// empty TimeInterval
	if ConfigPojo.TimeInterval == 0 {
		ConfigPojo.TimeInterval = 600
	}

	return ConfigPojo
}

func GetServiceCode(serviceName *string) string {
	serviceMap := map[string]string{
		"1":   "1", //移动
		"移动":  "1",
		"yd":  "1",
		"2":   "2", //联通
		"联通":  "2",
		"lt":  "2",
		"3":   "3", //电信
		"电信":  "3",
		"dx":  "3",
		"0":   "0", //校园网
		"校园网": "0",
		"教育网": "0",
		"edu": "0",
	}

	codeIdx := serviceMap[*serviceName] // code index
	codeMap := map[string]string{
		"1": "%E7%A7%BB%E5%8A%A8t",         //移动
		"2": "%E8%81%94%E9%80%9A",          //联通
		"3": "%E7%94%B5%E4%BF%A1",          //电信
		"0": "%E6%A0%A1%E5%9B%AD%E7%BD%91", //校园网
	}
	nameMap := map[string]string{
		"1": "移动",  //移动
		"2": "联通",  //联通
		"3": "电信",  //电信
		"0": "校园网", //校园网
	}
	*serviceName = nameMap[codeIdx] // update service name

	return codeMap[codeIdx]
}

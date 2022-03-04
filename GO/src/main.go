package main

import (
	"GO/src/pojo"
	utils2 "GO/src/utils"
	"strings"
	"time"
)

func main() {
	utils2.ReadConfig()
	resString, resCode := utils2.Get("http://www.google.cn/generate_204")
	for {
		for resCode != 204 {
			loginpageUrl := strings.Split(resString, "'")[1]
			loginUrl := strings.ReplaceAll(strings.Split(loginpageUrl, "?")[0], "index.jsp", "InterFace.do?method=login")
			queryString := strings.Split(loginpageUrl, "?")[1]
			queryString = strings.ReplaceAll(queryString, "&", "%2526")
			queryString = strings.ReplaceAll(queryString, "=", "%253D")
			config := utils2.ReadConfig()
			//service转换
			serviceMap := map[string]string{
				"1": "%E7%A7%BB%E5%8A%A8t",         //移动
				"2": "%E8%81%94%E9%80%9A",          //联通
				"3": "%E7%94%B5%E4%BF%A1",          //电信
				"4": "%E6%A0%A1%E5%9B%AD%E7%BD%91", //校园网
			}
			utils2.Post(loginUrl, pojo.UserData{
				UserId:      config.UserId,
				Password:    config.Password,
				Service:     serviceMap[config.Service],
				QueryString: queryString,
			})
			print("链接了一次网络\n")
			resString, resCode = utils2.Get("http://www.google.cn/generate_204")
			time.Sleep(1 * 100000000)
		}
		resString, resCode = utils2.Get("http://www.google.cn/generate_204")
		print("当前设备已经在线\n")
		time.Sleep(60 * 1000000000)
	}
}

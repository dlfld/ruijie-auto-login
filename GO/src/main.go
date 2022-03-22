package main

import (
	"fmt"
	"strings"
	"time"

	ruijielogger "GO/src/logger"
	pojo "GO/src/pojo"
	utils2 "GO/src/utils"
)

// type ConfigData struct {
// 	UserId       string `yaml:"UserId"`
// 	Password     string `yaml:"Password"`
// 	Service      string `yaml:"Service"`
// 	TimeInterval int    `yaml:"TimeInterval"`
// 	LogPath      string `yaml:"LogPath"`
// 	LogSaveDay   int    `yaml:"LogSaveDay"`
// 	LogClearDay  int    `yaml:"LogClearDay"`
// }

func main() {
	config := utils2.ReadConfig()
	logger := ruijielogger.NewRuijieLogger(config.LogPath, config.LogSaveDay, config.LogClearDay)

	// if config.TimeInterval < (60 * 3) {
	// 	config.TimeInterval = 60 * 3 // sleep 3 minutes
	// }

	for {
		resString, resCode := utils2.Get("http://www.google.cn/generate_204")
		for resCode != 204 {
			loginpageUrl := strings.Split(resString, "'")[1]
			loginUrl := strings.ReplaceAll(strings.Split(loginpageUrl, "?")[0], "index.jsp", "InterFace.do?method=login")
			queryString := strings.Split(loginpageUrl, "?")[1]
			queryString = strings.ReplaceAll(queryString, "&", "%2526")
			queryString = strings.ReplaceAll(queryString, "=", "%253D")

			//transformer config.server to Standard Server Name
			serverCode := utils2.GetServiceCode(&config.Server)

			utils2.Post(loginUrl, &pojo.UserData{
				UserId:      config.UserId,
				Password:    config.Password,
				Server:      serverCode,
				QueryString: queryString,
			})
			logger.Log(fmt.Sprintf("Try connect to %s with User %s", config.Server, config.UserId))
			resString, resCode = utils2.Get("http://www.google.cn/generate_204")
			time.Sleep(time.Duration(config.TimeInterval) * time.Second)
			logger.Log("Get below infos: ")
			logger.Log(resString)
			logger.Log(fmt.Sprintf("ResCode: %d", resCode))
		}
		logger.Log("Already online.")
		time.Sleep(time.Duration(config.TimeInterval) * time.Second)
	}
}

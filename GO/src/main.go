package main

import (
	"fmt"
	"strings"
	"time"

	ruijielogger "GO/src/logger"
	pojo "GO/src/pojo"
	configUtils "GO/src/service"
	utils2 "GO/src/utils"
)

func main() {
	//开始运行的时候检测是否还有配置文件，如果有的话就直接运行，如果没有的话就生成配置文件并提醒用户重新运行
	haveConfig := configUtils.DetectConfig()
	if !haveConfig {
		print("配置文件已生成，请填写配置之后重启该软件。\n")
		return
	}
	config := utils2.ReadConfig()
<<<<<<< HEAD
	logger := ruijielogger.NewRuijieLogger(config.LogPath, config.LogSaveDay, config.LogClearDay)
=======
	logger := ruijielogger.NewRuijieLogger(config.LogPath, config.LogClearDay)
	logger.Log("Start RuijieAL")
	logger.Log("User:" + config.UserId)
	logger.Log("Password:" + config.Password)
	logger.Log("LogPath:" + config.LogPath)

	if config.TimeInterval < (60 * 3) {
		config.TimeInterval = 60 * 3 // sleep 3 minutes
	}

>>>>>>> 7de7aa138e0ef970e55ae3c6a4dd6d94a8fd2887
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

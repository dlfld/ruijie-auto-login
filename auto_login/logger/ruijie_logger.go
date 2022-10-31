package logger

import (
	"fmt"
	"os"
	"strings"

	utils "auto_login/utils"

	go_logger "github.com/phachon/go-logger"
)

type RuijieLogger struct {
	TimeClear int               // clear logs
	LogPath   string            //log path
	Logger    *go_logger.Logger //go-logger
}

// NewRuijieLogger Init RuijieLogger
func NewRuijieLogger(logpath string,
	timeclear int) *RuijieLogger {

	// if !strings.Contains(logpath, ".log") {
	// 	logpath = logpath + ".log." + utils.GetNowTimeString()
	// }
	//logpath = logpath + "." + utils.GetNowTimeString()  // will

	logger := go_logger.NewLogger()
	logger.Detach("console")

	// console adapter config
	consoleConfig := &go_logger.ConsoleConfig{
		Color:      true, // Does the text display the color
		JsonFormat: true, // Whether or not formatted into a JSON string
		Format:     "",   // JsonFormat is false, logger message output to console format string
	}
	// add output to the console
	logger.Attach("console", go_logger.LOGGER_LEVEL_DEBUG, consoleConfig)

	fileConfig := &go_logger.FileConfig{
		Filename: logpath, // 日志输出文件名，不自动存在
		// 如果要将单独的日志分离为文件，请配置LealFrimeNem参数。
		// LevelFileName : map[int]string {
		//     logger.LoggerLevel("error"): "./error.log",    // Error 级别日志被写入 error .log 文件
		//     logger.LoggerLevel("info"): "./info.log",      // Info 级别日志被写入到 info.log 文件中
		//     logger.LoggerLevel("debug"): "./debug.log",    // Debug 级别日志被写入到 debug.log 文件中
		// },
		// MaxSize:    0,     // 文件最大值（KB），默认值0不限
		// MaxLine:    0,     // 文件最大行数，默认 0 不限制
		DateSlice:  "d",   // 文件根据日期切分， 支持 "Y" (年), "m" (月), "d" (日), "H" (时)
		JsonFormat: false, // 写入文件的数据是否 json 格式化
	}
	// 添加 file 为 logger 的一个输出
	err := logger.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileConfig)
	if err != nil {
		return nil
	}

	return &RuijieLogger{
		TimeClear: timeclear,
		LogPath:   logpath,
		Logger:    logger,
	}
}

// GetLoggerFilename logger defualt strategy, filename format as "ral_20220101.log"
func GetLoggerFilename(logpath string, str_time string) string {
	iPoint := strings.LastIndex(logpath, ".")

	filename := logpath
	fileExt := ""

	if iPoint != -1 {
		filename = logpath[:iPoint]
		fileExt = logpath[iPoint:]
	}

	filepath := fmt.Sprintf("%s_%s%s", filename, str_time, fileExt)
	return filepath
}

func IsExistsLog(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil || os.IsExist(err)
}

func (ruijieLogger *RuijieLogger) Log(info string) {
	strClear := utils.GetDaysAgoTimeString(ruijieLogger.TimeClear) // clear time
	// str_backup := utils.GetDaysAgoTimeString(ruijieLogger.TimeBackup) // backup time

	clearFilepath := GetLoggerFilename(ruijieLogger.LogPath, strClear)
	ruijieLogger.Logger.Info(fmt.Sprintf("Clear path: %s", clearFilepath))

	if ruijieLogger.TimeClear > 0 && IsExistsLog(clearFilepath) {
		// clear log
		err := os.Remove(clearFilepath)
		if err != nil {
			// 删除失败
			ruijieLogger.Logger.Error(fmt.Sprintf("Delete %s failed!", clearFilepath))
		} else {
			// 删除成功
			ruijieLogger.Logger.Info(fmt.Sprintf("Delete %s", clearFilepath))
		}
	}

	ruijieLogger.Logger.Info(info)
}

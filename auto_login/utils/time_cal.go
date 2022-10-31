package utils

import "time"

func GetNowTimeString() string {
	strTime := time.Now().Format("2006-01-02")
	return strTime
}

func GetDaysAgoTimeString(DaysAgo int) string {
	strClearTime := time.Now().AddDate(0, 0, -DaysAgo).Format("20060102") // get clearday string
	return strClearTime
}

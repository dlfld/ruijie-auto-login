package utils

import "time"

func GetNowTimeString() string {
	str_time := time.Now().Format("2006-01-02")
	return str_time
}

func GetDaysAgoTimeString(DaysAgo int) string {
	str_clear_time := time.Now().AddDate(0, 0, -DaysAgo).Format("20060102") // get clearday string
	return str_clear_time
}

package ssau

import (
	"strconv"
	"strings"
	"time"
)

func GetIdFromURL(url string) int64 {
	parts := strings.Split(url, "=")
	strID := parts[len(parts)-1]

	id, _ := strconv.ParseInt(strID, 10, 64)
	return id
}

func GetWeekday() int {
	now := time.Now()
	weekday := now.Weekday() - 1
	if weekday < 0 {
		weekday += 7
	}
	return int(weekday)
}

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

func GetWeekday(offsetDays int) int {
	offset := time.Duration(offsetDays) * 24 * time.Hour
	now := time.Now().Add(offset)
	weekday := now.Weekday() - 1
	if weekday < 0 {
		weekday += 7
	}
	return int(weekday)
}

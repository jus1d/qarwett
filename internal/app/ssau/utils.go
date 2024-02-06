package ssau

import (
	"strconv"
	"strings"
	"time"
	"unicode"
)

// GetIdFromURL parse link to the university's website, and grabs ID from there.
// E.g. https://ssau.ru/rasp?stafId=7819203 Grabbed ID will be: 7819203.
func GetIdFromURL(url string) int64 {
	parts := strings.Split(url, "=")
	strID := parts[len(parts)-1]

	id, _ := strconv.ParseInt(strID, 10, 64)
	return id
}

// GetWeekday returns today's weekday (if provided offset is 0), and today + offset (if provided offset != 0).
func GetWeekday(offsetDays int) int {
	offset := time.Duration(offsetDays) * 24 * time.Hour
	now := time.Now().Add(offset)
	weekday := now.Weekday() - 1
	if weekday < 0 {
		weekday += 7
	}
	// time.Weekday() returns weekday in week format where Sunday is the first weekday.
	// This lines transform weekday to format where first day of week is Monday.
	return int(weekday)
}

// Capitalize capitalizes first char of string.
func Capitalize(s string) string {
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

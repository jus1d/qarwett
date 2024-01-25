package ssau

import (
	"strconv"
	"strings"
)

func GetIdFromURL(url string) int64 {
	parts := strings.Split(url, "=")
	strID := parts[len(parts)-1]

	id, _ := strconv.ParseInt(strID, 10, 64)
	return id
}

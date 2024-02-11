package callback

import (
	"strconv"
	"strings"
)

func ExtractFromScheduleCallback(query string) (groupID int64, groupTitle string, week int, weekday int) {
	parts := strings.Split(query, ":")

	groupID, _ = strconv.ParseInt(parts[1], 10, 64)
	groupTitle = parts[2]
	week, _ = strconv.Atoi(parts[3])
	weekday, _ = strconv.Atoi(parts[4])

	return groupID, groupTitle, week, weekday
}

func ExtractFromScheduleTodayCallback(query string) (groupID int64, groupTitle string) {
	parts := strings.Split(query, ":")

	groupID, _ = strconv.ParseInt(parts[1], 10, 64)
	groupTitle = parts[2]

	return groupID, groupTitle
}

func ExtractFromFavouriteGroupCallback(query string) (groupID int64, groupTitle string) {
	parts := strings.Split(query, ":")

	groupID, _ = strconv.ParseInt(parts[1], 10, 64)
	groupTitle = parts[2]

	return groupID, groupTitle
}

func ExtractFromAddCalendarCallback(query string) (groupID int64, languageCode string) {
	parts := strings.Split(query, ":")

	groupID, _ = strconv.ParseInt(parts[1], 10, 64)
	languageCode = parts[2]

	return groupID, languageCode
}

func ExtractFromSetLanguageCallback(query string) (languageCode string) {
	parts := strings.Split(query, ":")

	languageCode = parts[1]

	return languageCode
}

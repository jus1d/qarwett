package callback

import "fmt"

func ApplyScheduleMask(groupID int64, groupTitle string, week int, weekday int) string {
	return fmt.Sprintf("schedule-daily:%d:%s:%d:%d", groupID, groupTitle, week, weekday)
}

func ApplyScheduleTodayMask(groupID int64, groupTitle string) string {
	return fmt.Sprintf("schedule-today:%d:%s", groupID, groupTitle)
}

func ApplyFavouriteGroupMask(groupID int64, groupTitle string) string {
	return fmt.Sprintf("favourite-group:%d:%s", groupID, groupTitle)
}

func ApplyAddCalendarMask(groupID int64, languageCode string) string {
	return fmt.Sprintf("add-calendar:%d:%s", groupID, languageCode)
}

func ApplySetLanguageMask(languageCode string) string {
	return fmt.Sprintf("set-language:%s", languageCode)
}

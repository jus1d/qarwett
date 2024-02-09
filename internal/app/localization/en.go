package localization

var EnglishLocalization = Locale{
	Messages: Messages{},
	Buttons: Buttons{
		Today:       "Today",
		Favourite:   "To Favourites",
		Cancel:      "Cancel",
		Approve:     "Approve",
		AddCalendar: "Add to Calendar",
	},
	languageCode: English,
}

func aboutEN(commit string) string {
	return ""
}

func announcementCheckEN(content string) string {
	return ""
}

func usersAmountEN(amount int) string {
	return ""
}

func freeDayEN(day int, month int) string {
	return ""
}

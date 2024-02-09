package localization

var RussianLocalization = Locale{
	Messages: Messages{},
	Buttons: Buttons{
		Today:       "Сегодня",
		Favourite:   "В Избранное",
		Cancel:      "Отмена",
		Approve:     "Подтвердить",
		AddCalendar: "Добавить в Календарь",
	},
	languageCode: Russian,
}

func aboutRU(commit string) string {
	return ""
}

func announcementCheckRU(content string) string {
	return ""
}

func usersAmountRU(amount int) string {
	return ""
}

func freeDayRU(day int, month int) string {
	return ""
}

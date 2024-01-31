package locale

func ScheduleCalendarName(languageCode string) string {
	switch languageCode {
	case RU:
		return "Расписание"
	default:
		return "University Schedule"
	}
}

func ScheduleSubgroup(languageCode string) string {
	switch languageCode {
	case RU:
		return "Подгруппа"
	default:
		return "Subgroup"
	}
}

func ScheduleIn(languageCode string) string {
	switch languageCode {
	case RU:
		return "в"
	default:
		return "in"
	}
}

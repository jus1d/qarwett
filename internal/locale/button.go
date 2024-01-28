package locale

func GetButtonToday(languageCode string) string {
	switch languageCode {
	case RU:
		return "Сегодня"
	default:
		return "Today"
	}
}

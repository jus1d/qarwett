package locale

func GetButtonToday(languageCode string) string {
	switch languageCode {
	case RU:
		return "Сегодня"
	default:
		return "Today"
	}
}

func GetButtonCancel(languageCode string) string {
	switch languageCode {
	case RU:
		return "Отмена"
	default:
		return "Cancel"
	}
}

func GetButtonApprove(languageCode string) string {
	switch languageCode {
	case RU:
		return "Подтвердить"
	default:
		return "Approve"
	}
}

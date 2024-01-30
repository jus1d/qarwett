package locale

func ButtonToday(languageCode string) string {
	switch languageCode {
	case RU:
		return "Сегодня"
	default:
		return "Today"
	}
}

func ButtonFavourite(languageCode string) string {
	switch languageCode {
	case RU:
		return "В избранное"
	default:
		return "Favourite"
	}
}

func ButtonCancel(languageCode string) string {
	switch languageCode {
	case RU:
		return "Отмена"
	default:
		return "Cancel"
	}
}

func ButtonApprove(languageCode string) string {
	switch languageCode {
	case RU:
		return "Подтвердить"
	default:
		return "Approve"
	}
}

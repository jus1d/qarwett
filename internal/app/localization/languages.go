package localization

type Language struct {
	Title        string
	LanguageCode string
}

var Languages = []Language{
	{
		Title:        "🇷🇺 Русский",
		LanguageCode: Russian,
	},
	{
		Title:        "🇺🇸 English",
		LanguageCode: English,
	},
}

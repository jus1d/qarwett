package locale

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	EN = "en"
	RU = "ru"
)

const unknownLanguageCodeMessage = "ERROR: unknown language code"

var PhrasesForFreeDayRU = []string{
	"<b>%d %s</b> - день бездельник",
	"<b>%d %s</b> даже будильник будет спать",
	"План на <b>%d %s</b> - не иметь планов",
	"<b>%d %s</b> можно и в Сочи рвануть",
	"<b>%d %s</b> разрешаю лениться",
	"<b>%d %s</b> - день исследований, исследуй куда постоянно деваются вторые носки",
	"Хоть <b>%d %s</b> и выходной, хватит на диване отлеживаться",
}

var PhrasesForFreeDayEN = []string{
	"<b>%d %s</b> - a day of idleness",
	"<b>%d %s</b> even the alarm clock will be sleeping",
	"Plan for <b>%d %s</b> - just have no plans",
	"<b>%d %s</b> is a good day to head to Sochi",
	"<b>%d %s</b> - I allow you to be lazy",
	"<b>%d %s</b> - a day of exploration, investigate where the second socks constantly disappear",
	"Although it's <b>%d %s</b> off, stop lying on the couch",
}

func GetPhraseGreeting(languageCode string) string {
	switch languageCode {
	case EN:
		return "<b>Hello, here you can take a quick look at your schedule <s>and go get some sleep</s></b>\n\n" +
			"👇Just type your group"
	case RU:
		return "<b>Привет, здесь ты сможешь быстро посмотреть свое расписание <s>и пойти отсыпаться</s></b>\n\n" +
			"👇Просто напиши свою группу"
	}
	return unknownLanguageCodeMessage
}

func GetPhraseNoGroupFound(languageCode string) string {
	switch languageCode {
	case EN:
		return "☹️There are no groups at your request"
	case RU:
		return "☹️По твоему запросу нет групп"
	}
	return unknownLanguageCodeMessage
}

func GetPhraseChooseGroup(languageCode string) string {
	switch languageCode {
	case EN:
		return "🤔<b>Choose a group</b>"
	case RU:
		return "🤔<b>Выбери группу</b>"
	}
	return unknownLanguageCodeMessage
}

func GetPhraseNoScheduleFound(languageCode string) string {
	switch languageCode {
	case EN:
		return "🚨<b>Can't found schedule!</b>"
	case RU:
		return "🚨<b>Не могу найти расписание!</b>"
	}
	return unknownLanguageCodeMessage
}

func GetPhraseNoChanges(languageCode string) string {
	switch languageCode {
	case EN:
		return "No changes"
	case RU:
		return "Изменений нет"
	}
	return unknownLanguageCodeMessage
}

func GetRandomPhraseForFreeDay(languageCode string, day int, month int) string {
	monthsRU := []string{"", "января", "февраля", "марта", "апреля", "мая", "июня", "июля", "августа", "сентября", "октября", "сентября", "декабря"}
	monthsEN := []string{"", "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	switch languageCode {
	case EN:
		return fmt.Sprintf(choice(PhrasesForFreeDayEN), day, monthsEN[month])
	case RU:
		return fmt.Sprintf(choice(PhrasesForFreeDayRU), day, monthsRU[month])
	}
	return unknownLanguageCodeMessage
}

func choice(arr []string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	idx := r.Intn(len(arr))
	return arr[idx]
}

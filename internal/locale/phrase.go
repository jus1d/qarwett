package locale

import (
	"fmt"
	"math/rand"
	"time"
)

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
	"<b>%d %s</b> is a perfect day to head to Sochi",
	"<b>%d %s</b> - I allow you to be lazy",
	"<b>%d %s</b> - a day of exploration, investigate where the second socks constantly disappear",
	"Although it's <b>%d %s</b> off, stop lying on the couch",
}

func GetPhraseGreeting(languageCode string) string {
	switch languageCode {
	case RU:
		return "<b>Привет, здесь ты сможешь быстро посмотреть свое расписание <s>и пойти отсыпаться</s></b>\n\n" +
			"👇Просто напиши свою группу"
	default:
		return "<b>Hello, here you can take a quick look at your schedule <s>and go get some sleep</s></b>\n\n" +
			"👇Just type your group"
	}
}

func GetPhraseNoGroupFound(languageCode string) string {
	switch languageCode {
	case RU:
		return "☹️По твоему запросу нет групп"
	default:
		return "☹️There are no groups at your request"
	}
}

func GetPhraseChooseGroup(languageCode string) string {
	switch languageCode {
	case RU:
		return "🤔<b>Выбери группу</b>"
	default:
		return "🤔<b>Choose a group</b>"
	}
}

func GetPhraseNoScheduleFound(languageCode string) string {
	switch languageCode {
	case RU:
		return "🚨<b>Не могу найти расписание!</b>"
	default:
		return "🚨<b>Can't found schedule!</b>"
	}
}

func GetPhraseNoChanges(languageCode string) string {
	switch languageCode {
	case RU:
		return "Изменений нет"
	default:
		return "No changes"
	}
}

func GetPhraseForFreeDay(languageCode string, day int, month int) string {
	monthsRU := []string{"", "января", "февраля", "марта", "апреля", "мая", "июня", "июля", "августа", "сентября", "октября", "сентября", "декабря"}
	monthsEN := []string{"", "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	switch languageCode {
	case RU:
		return fmt.Sprintf(choice(PhrasesForFreeDayRU), day, monthsRU[month])
	default:
		return fmt.Sprintf(choice(PhrasesForFreeDayEN), day, monthsEN[month])
	}
}

func choice(arr []string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	idx := r.Intn(len(arr))
	return arr[idx]
}

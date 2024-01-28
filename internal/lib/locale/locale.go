package locale

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var PhrasesForFreeDay = []string{
	"<b>{day} {month}</b> - день бездельник",
	"<b>{day} {month}</b> даже будильник будет спать",
	"План на <b>{day} {month}</b> - не иметь планов",
	"<b>{day} {month}</b> можно и в Сочи рвануть",
	"<b>{day} {month}</b> разрешаю лениться",
	"<b>{day} {month}</b> - день исследований, исследуй куда постоянно деваются вторые носки",
	"Хоть <b>{day} {month}</b> и выходной, хватит на диване отлеживаться",
}

func GetPhraseGreeting(languageCode string) string {
	return "<b>Привет, здесь ты сможешь быстро посмотреть свое расписание <s>и пойти отсыпаться</s></b>\n\n" +
		"👇Просто напиши свою группу"
}

func GetPhraseNoGroupFound(languageCode string) string {
	return "☹️По твоему запросу нет групп"
}

func GetPhraseChooseGroup(languageCode string) string {
	return "🤔<b>Выбери группу</b>"
}

func GetPhraseNoScheduleFound(languageCode string) string {
	return "🚨<b>Не могу найти расписание!</b>"
}

func GetPhraseNoChanges(languageCode string) string {
	return "Изменений нет"
}

func GetRandomPhraseForFreeDay(languageCode string, day int, month string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	idx := r.Intn(len(PhrasesForFreeDay))
	s := PhrasesForFreeDay[idx]
	s = strings.Replace(s, "{day}", strconv.Itoa(day), 1)
	s = strings.Replace(s, "{month}", month, 1)
	return s
}

package locale

import (
	"fmt"
	"math/rand"
	"time"
)

var PhrasesForFreeDay = []string{
	"<b>%d %s</b> - день бездельник",
	"<b>%d %s</b> даже будильник будет спать",
	"План на <b>%d %s</b> - не иметь планов",
	"<b>%d %s</b> можно и в Сочи рвануть",
	"<b>%d %s</b> разрешаю лениться",
	"<b>%d %s</b> - день исследований, исследуй куда постоянно деваются вторые носки",
	"Хоть <b>%d %s</b> и выходной, хватит на диване отлеживаться",
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

	return fmt.Sprintf(s, day, month)
}

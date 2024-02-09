package localization

import (
	"fmt"
	"qarwett/pkg/random"
)

var RussianLocalization = Locale{
	Message: Messages{
		Greeting: "<b>Привет, здесь ты сможешь быстро посмотреть свое расписание <s>и пойти отсыпаться</s></b>\n\n" +
			"👇Просто напиши свою группу",
		RequestAnnouncement:        "Отправьте сообщение для объявления",
		CantStartAnnouncement:      "Ошибка при создании объявления!",
		AnnouncementCompleted:      "<b>Объявление успешно разослано все пользователям</b>",
		CantFoundYourGroup:         "☹️Не могу найти твою группу\n\nЧтобы использовать эту команду, добавь свою группу <b>в избранное</b>",
		NoGroupFound:               "☹️По твоему запросу нет групп",
		ChooseGroup:                "🤔<b>Выбери группу</b>",
		NoScheduleFound:            "🚨<b>Не могу найти расписание!</b>",
		NoChanges:                  "Изменений нет",
		Error:                      "Ошибка",
		Success:                    "Успешно",
		Cancelled:                  "Отменено",
		FailedToCancel:             "Ошибка при отмене",
		AnnouncementMessageIsEmpty: "Сообщение для объявления не найдено. Попробуйте снова",
		UseRestart:                 "Требуется перезагрузка бота -> <b>/start</b>",
		AdminCommands: "<b>Список админ комманд:</b>\n\n" +
			"<b>/announce</b> - Отправить объявление всем пользователям бота\n" +
			"<b>/users</b> - Посмотреть количество пользователей",
		languageCode: Russian,
	},
	Button: Buttons{
		Today:       "Сегодня",
		Favourite:   "В Избранное",
		Cancel:      "Отмена",
		Approve:     "Подтвердить",
		AddCalendar: "Добавить в Календарь",
	},
	languageCode: Russian,
}

var phrasesForFreeDayRU = []string{
	"<b>%d %s</b> - день бездельник",
	"<b>%d %s</b> даже будильник будет спать",
	"План на <b>%d %s</b> - не иметь планов",
	"<b>%d %s</b> можно и в Сочи рвануть",
	"<b>%d %s</b> разрешаю лениться",
	"<b>%d %s</b> - день исследований, исследуй куда постоянно деваются вторые носки",
	"Хоть <b>%d %s</b> и выходной, хватит на диване отлеживаться",
}

func aboutRU(commit string) string {
	if commit == "" {
		commit = "xx"
	}

	return "<b>qarweTT</b> - быстро глянуть на расписание, и обратно спать\n\n" +
		"По всем вопросам, можете обращаться сюда <b>@jus1d</b>\n\n" +
		"Текущая версия сборки: <b>" + commit + "</b>"
}

func announcementCheckRU(content string) string {
	return "<b>Ваше объявление:</b>\n\n" + content + "\n\n" +
		"Все ле корректно?"
}

func usersAmountRU(amount int) string {
	return fmt.Sprintf("<b>Всего пользователей:</b> %d", amount)
}

func freeDayRU(day int, month int) string {
	monthsRU := []string{"", "января", "февраля", "марта", "апреля", "мая", "июня", "июля", "августа", "сентября", "октября", "сентября", "декабря"}

	return fmt.Sprintf(random.Choice(phrasesForFreeDayRU), day, monthsRU[month])
}

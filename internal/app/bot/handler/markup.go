package handler

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"qarwett/internal/app/bot/callback"
	"qarwett/internal/app/localization"
	"qarwett/internal/app/ssau"
)

func GetMarkupCancel(languageCode string) *telegram.InlineKeyboardMarkup {
	locale := localization.Get(languageCode)

	markup := telegram.NewInlineKeyboardMarkup(
		telegram.NewInlineKeyboardRow(
			telegram.NewInlineKeyboardButtonData(locale.Button.Cancel, "cancel"),
		),
	)
	return &markup
}

func GetMarkupCheckAnnouncement(languageCode string) *telegram.InlineKeyboardMarkup {
	locale := localization.Get(languageCode)

	markup := telegram.NewInlineKeyboardMarkup(
		telegram.NewInlineKeyboardRow(
			telegram.NewInlineKeyboardButtonData(locale.Button.Approve, "approve-announcement"),
			telegram.NewInlineKeyboardButtonData(locale.Button.Cancel, "cancel"),
		),
	)
	return &markup
}

func GetScheduleNavigationMarkup(languageCode string, groupID int64, groupTitle string, week int, weekday int, addFavourite bool) *telegram.InlineKeyboardMarkup {
	locale := localization.Get(languageCode)

	prevWeek := week
	prevWeekday := weekday - 1
	nextWeek := week
	nextWeekday := weekday + 1
	if prevWeekday == -1 {
		prevWeekday = 6
		prevWeek = week - 1
	}
	if nextWeekday == 7 {
		nextWeekday = 0
		nextWeek = week + 1
	}
	queryLeft := callback.ApplyScheduleMask(groupID, groupTitle, prevWeek, prevWeekday)
	queryUpdate := callback.ApplyScheduleMask(groupID, groupTitle, week, weekday)
	queryRight := callback.ApplyScheduleMask(groupID, groupTitle, nextWeek, nextWeekday)

	rows := make([][]telegram.InlineKeyboardButton, 0)

	rows = append(rows, telegram.NewInlineKeyboardRow(
		telegram.NewInlineKeyboardButtonData("«", queryLeft),
		telegram.NewInlineKeyboardButtonData("↻", queryUpdate),
		telegram.NewInlineKeyboardButtonData("»", queryRight),
	))

	rows = append(rows, telegram.NewInlineKeyboardRow(
		telegram.NewInlineKeyboardButtonData(locale.Button.Today, callback.ApplyScheduleTodayMask(groupID, groupTitle)),
	))

	if addFavourite {
		rows = append(rows, telegram.NewInlineKeyboardRow(
			telegram.NewInlineKeyboardButtonData(locale.Button.Favourite, callback.ApplyFavouriteGroupMask(groupID, groupTitle)),
		))
	}

	rows = append(rows, telegram.NewInlineKeyboardRow(
		telegram.NewInlineKeyboardButtonData(locale.Button.AddCalendar, callback.ApplyAddCalendarMask(groupID, languageCode)),
	))

	markup := telegram.NewInlineKeyboardMarkup(rows...)

	return &markup
}

func GetMarkupFromGroupList(groups []ssau.SearchGroupResponse) *telegram.InlineKeyboardMarkup {
	length := len(groups)
	for length%3 != 0 {
		length++
	}
	rows := make([][]telegram.InlineKeyboardButton, length/3)

	for i := 0; i < len(groups); i += 3 {
		buttons := make([]telegram.InlineKeyboardButton, 0)
		for j := 0; j < 3 && i+j < len(groups); j++ {
			group := groups[i+j]
			query := callback.ApplyScheduleMask(group.ID, group.Title, 0, ssau.GetWeekday(0))
			buttons = append(buttons, telegram.NewInlineKeyboardButtonData(group.Title, query))
		}
		rows[i/3] = telegram.NewInlineKeyboardRow(buttons...)
	}

	var markup telegram.InlineKeyboardMarkup
	markup.InlineKeyboard = rows

	return &markup
}

func GetLanguagesMarkup() *telegram.InlineKeyboardMarkup {
	rows := make([][]telegram.InlineKeyboardButton, 0)

	for _, language := range localization.Languages {
		rows = append(rows, telegram.NewInlineKeyboardRow(
			telegram.NewInlineKeyboardButtonData(language.Title, callback.ApplySetLanguageMask(language.LanguageCode)),
		))
	}

	markup := telegram.NewInlineKeyboardMarkup(rows...)

	return &markup
}

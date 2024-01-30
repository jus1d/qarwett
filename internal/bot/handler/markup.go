package handler

import (
	"fmt"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"qarwett/internal/locale"
	"qarwett/internal/ssau"
)

func GetMarkupCancel(languageCode string) *telegram.InlineKeyboardMarkup {
	markup := telegram.NewInlineKeyboardMarkup(
		telegram.NewInlineKeyboardRow(
			telegram.NewInlineKeyboardButtonData(locale.ButtonCancel(languageCode), "cancel"),
		),
	)
	return &markup
}

func GetMarkupCheckAnnouncement(languageCode string) *telegram.InlineKeyboardMarkup {
	markup := telegram.NewInlineKeyboardMarkup(
		telegram.NewInlineKeyboardRow(
			telegram.NewInlineKeyboardButtonData(locale.ButtonApprove(languageCode), "approve-announcement"),
			telegram.NewInlineKeyboardButtonData(locale.ButtonCancel(languageCode), "cancel"),
		),
	)
	return &markup
}

func GetScheduleNavigationMarkup(groupID int64, groupTitle string, week int, weekday int, addFavourite bool) *telegram.InlineKeyboardMarkup {
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
	queryLeft := ApplyScheduleMask(groupID, groupTitle, prevWeek, prevWeekday)
	queryUpdate := ApplyScheduleMask(groupID, groupTitle, week, weekday)
	queryRight := ApplyScheduleMask(groupID, groupTitle, nextWeek, nextWeekday)

	rows := make([][]telegram.InlineKeyboardButton, 0)

	rows = append(rows, telegram.NewInlineKeyboardRow(
		telegram.NewInlineKeyboardButtonData("«", queryLeft),
		telegram.NewInlineKeyboardButtonData("⟳", queryUpdate),
		telegram.NewInlineKeyboardButtonData("»", queryRight),
	))

	rows = append(rows, telegram.NewInlineKeyboardRow(
		telegram.NewInlineKeyboardButtonData(locale.ButtonToday(locale.RU), ApplyScheduleTodayMask(groupID, groupTitle)),
	))

	if addFavourite {
		rows = append(rows, telegram.NewInlineKeyboardRow(
			telegram.NewInlineKeyboardButtonData(locale.ButtonFavourite(locale.RU), ApplyFavouriteGroupMask(groupID, groupTitle)),
		))
	}

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
			query := ApplyScheduleMask(group.ID, group.Title, 0, ssau.GetWeekday(0))
			buttons = append(buttons, telegram.NewInlineKeyboardButtonData(group.Title, query))
		}
		rows[i/3] = telegram.NewInlineKeyboardRow(buttons...)
	}

	var markup telegram.InlineKeyboardMarkup
	markup.InlineKeyboard = rows

	return &markup
}

func ApplyScheduleMask(groupID int64, groupTitle string, week int, weekday int) string {
	return fmt.Sprintf("schedule-daily:%d:%s:%d:%d", groupID, groupTitle, week, weekday)
}

func ApplyScheduleTodayMask(groupID int64, groupTitle string) string {
	return fmt.Sprintf("schedule-today:%d:%s", groupID, groupTitle)
}

func ApplyFavouriteGroupMask(groupID int64, groupTitle string) string {
	return fmt.Sprintf("favourite-group:%d:%s", groupID, groupTitle)
}

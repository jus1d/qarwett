package handler

import (
	"fmt"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"qarwett/internal/ssau"
)

func GetScheduleNavigationMarkup(groupID int64, week int, weekday int) telegram.InlineKeyboardMarkup {
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
	queryLeft := ApplyScheduleMask(groupID, prevWeek, prevWeekday)
	queryUpdate := ApplyScheduleMask(groupID, week, weekday)
	queryRight := ApplyScheduleMask(groupID, nextWeek, nextWeekday)

	return telegram.NewInlineKeyboardMarkup(
		telegram.NewInlineKeyboardRow(
			telegram.NewInlineKeyboardButtonData("«", queryLeft),
			telegram.NewInlineKeyboardButtonData("⟳", queryUpdate),
			telegram.NewInlineKeyboardButtonData("»", queryRight),
		),
	)
}

func GetMarkupFromGroupList(groups []ssau.SearchGroupResponse) telegram.InlineKeyboardMarkup {
	length := len(groups)
	for length%3 != 0 {
		length++
	}
	rows := make([][]telegram.InlineKeyboardButton, length/3)

	for i := 0; i < len(groups); i += 3 {
		buttons := make([]telegram.InlineKeyboardButton, 0)
		for j := 0; j < 3 && i+j < len(groups); j++ {
			group := groups[i+j]
			query := ApplyScheduleMask(group.ID, 0, ssau.GetWeekday(0))
			buttons = append(buttons, telegram.NewInlineKeyboardButtonData(group.Title, query))
		}
		rows[i/3] = telegram.NewInlineKeyboardRow(buttons...)
	}

	var markup telegram.InlineKeyboardMarkup
	markup.InlineKeyboard = rows

	return markup
}

func ApplyScheduleMask(groupID int64, week int, weekday int) string {
	return fmt.Sprintf("schedule-daily:%d:%d:%d", groupID, week, weekday)
}

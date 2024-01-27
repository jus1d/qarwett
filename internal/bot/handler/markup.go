package handler

import (
	"fmt"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"qarwett/internal/ssau"
)

func GetScheduleNavigationMarkup(groupID int64, offset int) telegram.InlineKeyboardMarkup { // Maybe add offset limit
	queryLeft := fmt.Sprintf("schedule:%d:%d", groupID, offset-1)
	queryRight := fmt.Sprintf("schedule:%d:%d", groupID, offset+1)
	queryUpdate := fmt.Sprintf("schedule:%d:%d", groupID, offset)
	return telegram.NewInlineKeyboardMarkup(
		telegram.NewInlineKeyboardRow(
			telegram.NewInlineKeyboardButtonData("«", queryLeft),
			telegram.NewInlineKeyboardButtonData("⟳", queryUpdate),
			telegram.NewInlineKeyboardButtonData("»", queryRight),
		),
	)
}

// TODO(#9): Change callback format
// Move out of this format: `schedule:groupID:offset` to this: `schedule:groupID:week:weekday`

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
			query := fmt.Sprintf("schedule:%d:0", group.ID)
			buttons = append(buttons, telegram.NewInlineKeyboardButtonData(group.Title, query))
		}
		rows[i/3] = telegram.NewInlineKeyboardRow(buttons...)
	}

	var markup telegram.InlineKeyboardMarkup
	markup.InlineKeyboard = rows

	return markup
}

package handler

import (
	"fmt"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"qarwett/internal/schedule"
)

func GetScheduleNavigationMarkup(groupID int64, offset int) telegram.InlineKeyboardMarkup { // Maybe add offset limit
	callbackLeft := fmt.Sprintf("schedule:%d:%d", groupID, offset-1)
	callbackRight := fmt.Sprintf("schedule:%d:%d", groupID, offset+1)
	return telegram.NewInlineKeyboardMarkup(
		telegram.NewInlineKeyboardRow(
			telegram.NewInlineKeyboardButtonData("«", callbackLeft),
			telegram.NewInlineKeyboardButtonData("»", callbackRight),
		),
	)
}

func GetMarkupFromGroupList(groups []schedule.Group) telegram.InlineKeyboardMarkup {
	return telegram.NewInlineKeyboardMarkup()
}

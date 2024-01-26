package handler

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"qarwett/internal/lib/logger/sl"
	"qarwett/internal/schedule"
	"qarwett/internal/ssau"
	"strconv"
	"strings"
)

func (h *Handler) OnCallbackSchedule(u telegram.Update) {
	author := u.CallbackQuery.From

	log := h.log.With(
		slog.String("op", "handler.OnCallbackSchedule"),
		slog.String("username", author.UserName),
		slog.String("id", strconv.FormatInt(author.ID, 10)),
	)

	query := u.CallbackData()
	log.Debug("Callback handled", slog.String("query", query))

	parts := strings.Split(query, ":")
	groupID, _ := strconv.ParseInt(parts[1], 10, 64)
	offset, _ := strconv.Atoi(parts[2])

	doc, err := ssau.GetScheduleDocument(groupID, 0)
	if err != nil {
		_, err = h.SendTextMessage(author.ID, "Can't get a schedule. Sorry!", nil)
		if err != nil {
			log.Error("Failed to send message")
			return
		}
	}
	timetable := ssau.Parse(doc)
	weekday := ssau.GetWeekday(offset)

	content := schedule.ParseScheduleToMessageTextWithHTML(timetable[weekday])

	_, err = h.EditMessageText(u.CallbackQuery.Message, content, GetScheduleNavigationMarkup(groupID, offset))
	if err != nil {
		log.Error("Failed to send message", sl.Err(err))
		return
	}
}

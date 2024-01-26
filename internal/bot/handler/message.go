package handler

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"qarwett/internal/lib/logger/sl"
	"qarwett/internal/schedule"
	"qarwett/internal/ssau"
	"strconv"
)

func (h *Handler) HandleMessage(u telegram.Update) {
	author := u.Message.From

	log := h.log.With(
		slog.String("op", "handler.HandleMessage"),
		slog.String("username", author.UserName),
		slog.String("id", strconv.FormatInt(author.ID, 10)),
	)

	log.Debug("Message handled", slog.String("content", u.Message.Text))

	query := u.Message.Text

	groups, err := ssau.GetGroupBySearchQuery(query)
	if len(groups) == 0 || err != nil {
		_, err := h.SendTextMessage(author.ID, "Can't found group '"+query+"'.", nil)
		if err != nil {
			log.Error("Failed to send message")
		}
		return
	}

	// TODO(#3): Create button with different groups, to get user an ability to choose
	group := groups[0]
	doc, err := ssau.GetScheduleDocument(group.ID, 25)
	if err != nil {
		_, err = h.SendTextMessage(author.ID, "Can't get a schedule. Sorry!", nil)
		if err != nil {
			log.Error("Failed to send message")
			return
		}
	}
	timetable := ssau.Parse(doc)
	weekday := ssau.GetWeekday(0)

	content := schedule.ParseScheduleToMessageTextWithHTML(timetable[weekday])

	_, err = h.SendTextMessage(author.ID, content, GetScheduleNavigationMarkup(group.ID, 0))
	if err != nil {
		log.Error("Failed to send message", sl.Err(err))
		return
	}
}

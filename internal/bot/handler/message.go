package handler

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"qarwett/internal/lib/logger/sl"
	"qarwett/internal/locale"
	"qarwett/internal/schedule"
	"qarwett/internal/ssau"
	"strconv"
)

func (h *Handler) OnNewMessage(u telegram.Update) {
	author := u.Message.From

	log := h.log.With(
		slog.String("op", "handler.OnNewMessage"),
		slog.String("username", author.UserName),
		slog.String("id", strconv.FormatInt(author.ID, 10)),
	)

	log.Debug("Message handled", slog.String("content", u.Message.Text))

	query := u.Message.Text

	groups, err := ssau.GetGroupBySearchQuery(query)
	if len(groups) == 0 || err != nil {
		_, err = h.SendTextMessage(author.ID, locale.GetPhraseNoGroupFound(locale.RU), nil)
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
		}
		return
	}

	if len(groups) > 1 {
		markup := GetMarkupFromGroupList(groups)
		_, err = h.SendTextMessage(author.ID, locale.GetPhraseChooseGroup(locale.RU), markup)
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
		}
		return
	}

	group := groups[0]
	doc, err := ssau.GetScheduleDocument(group.ID, 0)
	if err != nil {
		_, err = h.SendTextMessage(author.ID, locale.GetPhraseNoScheduleFound(locale.RU), nil)
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
			return
		}
	}
	timetable, week := ssau.Parse(doc)

	weekday := ssau.GetWeekday(0)
	content := schedule.ParseScheduleToMessageTextWithHTML(schedule.Day{
		Date:  timetable.StartDate.AddDate(0, 0, weekday),
		Pairs: timetable.Pairs[weekday],
	})

	_, err = h.SendTextMessage(author.ID, content, GetScheduleNavigationMarkup(group.ID, week, weekday))
	if err != nil {
		log.Error("Failed to send message", sl.Err(err))
		return
	}
}

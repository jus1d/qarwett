package handler

import (
	"errors"
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

	parts := strings.Split(query, ":") // {"schedule", groupID, week, weekday}
	groupID, _ := strconv.ParseInt(parts[1], 10, 64)
	week, _ := strconv.Atoi(parts[2])
	weekday, _ := strconv.Atoi(parts[3])

	doc, err := ssau.GetScheduleDocument(groupID, week)
	if err != nil {
		_, err = h.SendTextMessage(author.ID, "Can't get a schedule. Sorry!", nil)
		if err != nil {
			log.Error("Failed to send message")
			return
		}
	}
	timetable, week := ssau.Parse(doc)

	content := schedule.ParseScheduleToMessageTextWithHTML(schedule.Day{
		Date:  timetable.StartDate.AddDate(0, 0, weekday),
		Pairs: timetable.Pairs[weekday],
	})

	_, err = h.EditMessageText(u.CallbackQuery.Message, content, GetScheduleNavigationMarkup(groupID, week, weekday))
	if errors.Is(err, ErrNoChanges) {
		callback := telegram.NewCallback(u.CallbackQuery.ID, "Изменений нет")
		_, err = h.bot.Request(callback)
		if err != nil {
			log.Error("Failed to send callback", sl.Err(err))
		}
	}
	if err != nil {
		log.Error("Failed to edit message", sl.Err(err))
		return
	}
}

func (h *Handler) OnCallbackScheduleToday(u telegram.Update) {
	author := u.CallbackQuery.From

	log := h.log.With(
		slog.String("op", "handler.OnCallbackSchedule"),
		slog.String("username", author.UserName),
		slog.String("id", strconv.FormatInt(author.ID, 10)),
	)

	query := u.CallbackData()
	log.Debug("Callback handled", slog.String("query", query))

	parts := strings.Split(query, ":") // {"schedule-today", groupID}
	groupID, _ := strconv.ParseInt(parts[1], 10, 64)
	weekday := ssau.GetWeekday(0)

	doc, err := ssau.GetScheduleDocument(groupID, 0)
	if err != nil {
		_, err = h.SendTextMessage(author.ID, "Can't get a schedule. Sorry!", nil)
		if err != nil {
			log.Error("Failed to send message")
			return
		}
	}
	timetable, week := ssau.Parse(doc)

	content := schedule.ParseScheduleToMessageTextWithHTML(schedule.Day{
		Date:  timetable.StartDate.AddDate(0, 0, weekday),
		Pairs: timetable.Pairs[weekday],
	})

	_, err = h.EditMessageText(u.CallbackQuery.Message, content, GetScheduleNavigationMarkup(groupID, week, weekday))
	if errors.Is(err, ErrNoChanges) {
		callback := telegram.NewCallback(u.CallbackQuery.ID, "Изменений нет")
		_, err = h.bot.Request(callback)
		if err != nil {
			log.Error("Failed to send callback", sl.Err(err))
		}
	}
	if err != nil {
		log.Error("Failed to edit message", sl.Err(err))
		return
	}
}

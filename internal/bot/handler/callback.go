package handler

import (
	"errors"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"qarwett/internal/lib/logger/sl"
	"qarwett/internal/locale"
	"qarwett/internal/schedule"
	"qarwett/internal/ssau"
	"qarwett/internal/storage/postgres"
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
		_, err = h.SendTextMessage(author.ID, locale.GetPhraseNoScheduleFound(locale.RU), nil)
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
		callback := telegram.NewCallback(u.CallbackQuery.ID, locale.GetPhraseNoChanges(locale.RU))
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
		slog.String("op", "handler.OnCallbackScheduleToday"),
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
		_, err = h.SendTextMessage(author.ID, locale.GetPhraseNoScheduleFound(locale.RU), nil)
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
		callback := telegram.NewCallback(u.CallbackQuery.ID, locale.GetPhraseNoChanges(locale.RU))
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

func (h *Handler) OnCallbackCancel(u telegram.Update) {
	author := u.CallbackQuery.From

	log := h.log.With(
		slog.String("op", "handler.OnCallbackCancel"),
		slog.String("username", author.UserName),
		slog.String("id", strconv.FormatInt(author.ID, 10)),
	)

	query := u.CallbackData()
	log.Debug("Callback handled", slog.String("query", query))

	err := h.storage.UpdateUserStage(author.ID, postgres.StageNone)
	if err != nil {
		log.Error("Failed to update user's stage", sl.Err(err))
		callback := telegram.NewCallback(u.CallbackQuery.ID, locale.GetPhraseFailedToCancel(locale.RU))
		_, err = h.bot.Request(callback)
		if err != nil {
			log.Error("Failed to send callback", sl.Err(err))
		}
	}

	callback := telegram.NewCallback(u.CallbackQuery.ID, locale.GetPhraseCancelled(locale.RU))
	_, err = h.bot.Request(callback)
	if err != nil {
		log.Error("Failed to send callback", sl.Err(err))
	}

	c := telegram.NewDeleteMessage(author.ID, u.CallbackQuery.Message.MessageID)
	_, _ = h.bot.Send(c)
}

func (h *Handler) OnCallbackAnnouncementApprove(u telegram.Update) {
	author := u.CallbackQuery.From

	log := h.log.With(
		slog.String("op", "handler.OnCallbackAnnouncementApprove"),
		slog.String("username", author.UserName),
		slog.String("id", strconv.FormatInt(author.ID, 10)),
	)

	query := u.CallbackData()
	log.Debug("Callback handled", slog.String("query", query))

	content, exists := h.storage.GetAnnouncementMessage(author.ID)
	if !exists {
		log.Error("Content for announcement doesn't exists")
		callback := telegram.NewCallback(u.CallbackQuery.ID, locale.GetPhraseEmptyAnnouncementMessage(locale.RU))
		_, err := h.bot.Request(callback)
		if err != nil {
			log.Error("Failed to send callback", sl.Err(err))
		}
	}

	users, err := h.storage.GetAllUsers()
	if err != nil {
		log.Error("Failed to get all users", sl.Err(err))
		callback := telegram.NewCallback(u.CallbackQuery.ID, locale.GetPhraseCantStartAnnouncement(locale.RU))
		_, err = h.bot.Request(callback)
		if err != nil {
			log.Error("Failed to send callback", sl.Err(err))
		}
	}

	for _, user := range users {
		if user.TelegramID == author.ID {
			continue
		}
		_, err = h.SendTextMessage(user.TelegramID, content, nil)
		if err != nil {
			log.Error("Failed to send an announcement message", sl.Err(err), slog.Int64("recipientID", user.TelegramID))
		}
	}

	_, err = h.EditMessageText(u.CallbackQuery.Message, locale.GetPhraseAnnouncementCompleted(locale.RU), nil)
	if err != nil {
		log.Error("Failed to send message", sl.Err(err))
	}
}

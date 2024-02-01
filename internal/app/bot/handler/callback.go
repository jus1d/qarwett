package handler

import (
	"errors"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	locale2 "qarwett/internal/app/locale"
	"qarwett/internal/app/schedule"
	"qarwett/internal/app/ssau"
	"qarwett/internal/storage/postgres"
	"qarwett/pkg/logger/sl"
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

	parts := strings.Split(query, ":") // {"schedule", groupID, groupTitle, week, weekday}
	groupID, _ := strconv.ParseInt(parts[1], 10, 64)
	groupTitle := parts[2]
	week, _ := strconv.Atoi(parts[3])
	weekday, _ := strconv.Atoi(parts[4])

	doc, err := ssau.GetScheduleDocument(groupID, week)
	if err != nil {
		_, err = h.SendTextMessage(author.ID, locale2.PhraseNoScheduleFound(locale2.RU), nil)
		if err != nil {
			log.Error("Failed to send message")
			return
		}
	}
	timetable, week := ssau.Parse(doc)

	user, err := h.storage.GetUserByTelegramID(author.ID)
	if err != nil {
		log.Error("Failed to get user from storage", sl.Err(err))
	}

	favouriteButton := err == nil && user.LinkedGroupID != groupID

	content := schedule.ParseScheduleToMessageTextWithHTML(groupID, groupTitle, schedule.Day{
		Date:  timetable.StartDate.AddDate(0, 0, weekday),
		Pairs: timetable.Pairs[weekday],
	})

	_, err = h.EditMessageText(u.CallbackQuery.Message, content,
		GetScheduleNavigationMarkup(groupID, groupTitle, week, weekday, favouriteButton))
	if errors.Is(err, ErrNoChanges) {
		callback := telegram.NewCallback(u.CallbackQuery.ID, locale2.PhraseNoChanges(locale2.RU))
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

func (h *Handler) OnCallbackFavouriteGroup(u telegram.Update) {
	author := u.CallbackQuery.From

	log := h.log.With(
		slog.String("op", "handler.OnCallbackFavouriteGroup"),
		slog.String("username", author.UserName),
		slog.String("id", strconv.FormatInt(author.ID, 10)),
	)

	query := u.CallbackData()
	log.Debug("Callback handled", slog.String("query", query))

	parts := strings.Split(query, ":") // {"favourite-group", groupID}
	groupID, _ := strconv.ParseInt(parts[1], 10, 64)
	groupTitle := parts[2]

	err := h.storage.UpdateUserLinkedGroup(author.ID, groupID, groupTitle)
	if err != nil {
		log.Error("Failed to update user's group", sl.Err(err))
		callback := telegram.NewCallback(u.CallbackQuery.ID, locale2.PhraseError(locale2.RU))
		_, err = h.bot.Request(callback)
		if err != nil {
			log.Error("Failed to send callback", sl.Err(err))
		}
	}

	callback := telegram.NewCallback(u.CallbackQuery.ID, locale2.PhraseSuccess(locale2.RU))
	_, err = h.bot.Request(callback)
	if err != nil {
		log.Error("Failed to send callback", sl.Err(err))
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
	groupTitle := parts[2]
	weekday := ssau.GetWeekday(0)

	doc, err := ssau.GetScheduleDocument(groupID, 0)
	if err != nil {
		_, err = h.SendTextMessage(author.ID, locale2.PhraseNoScheduleFound(locale2.RU), nil)
		if err != nil {
			log.Error("Failed to send message")
			return
		}
	}
	timetable, week := ssau.Parse(doc)

	user, err := h.storage.GetUserByTelegramID(author.ID)
	if err != nil {
		log.Error("Failed to get user from storage", sl.Err(err))
	}

	favouriteButton := err == nil && user.LinkedGroupID != groupID

	content := schedule.ParseScheduleToMessageTextWithHTML(groupID, groupTitle, schedule.Day{
		Date:  timetable.StartDate.AddDate(0, 0, weekday),
		Pairs: timetable.Pairs[weekday],
	})

	_, err = h.EditMessageText(u.CallbackQuery.Message, content,
		GetScheduleNavigationMarkup(groupID, groupTitle, week, weekday, favouriteButton))
	if errors.Is(err, ErrNoChanges) {
		callback := telegram.NewCallback(u.CallbackQuery.ID, locale2.PhraseNoChanges(locale2.RU))
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
		callback := telegram.NewCallback(u.CallbackQuery.ID, locale2.PhraseFailedToCancel(locale2.RU))
		_, err = h.bot.Request(callback)
		if err != nil {
			log.Error("Failed to send callback", sl.Err(err))
		}
	}

	callback := telegram.NewCallback(u.CallbackQuery.ID, locale2.PhraseCancelled(locale2.RU))
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
		callback := telegram.NewCallback(u.CallbackQuery.ID, locale2.PhraseEmptyAnnouncementMessage(locale2.RU))
		_, err := h.bot.Request(callback)
		if err != nil {
			log.Error("Failed to send callback", sl.Err(err))
		}
	}

	users, err := h.storage.GetAllUsers()
	if err != nil {
		log.Error("Failed to get all users", sl.Err(err))
		callback := telegram.NewCallback(u.CallbackQuery.ID, locale2.PhraseCantStartAnnouncement(locale2.RU))
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

	_, err = h.EditMessageText(u.CallbackQuery.Message, locale2.PhraseAnnouncementCompleted(locale2.RU), nil)
	if err != nil {
		log.Error("Failed to send message", sl.Err(err))
	}
}
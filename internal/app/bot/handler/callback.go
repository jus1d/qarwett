package handler

import (
	"errors"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"os"
	"qarwett/internal/app/bot/callback"
	"qarwett/internal/app/icalendar"
	"qarwett/internal/app/localization"
	"qarwett/internal/app/schedule"
	"qarwett/internal/app/ssau"
	"qarwett/internal/storage/postgres"
	"qarwett/pkg/logger/sl"
	"strconv"
)

func (h *Handler) OnCallbackSchedule(u telegram.Update) {
	author := u.CallbackQuery.From

	log := h.log.With(
		slog.String("op", "handler.OnCallbackSchedule"),
		slog.String("username", author.UserName),
		slog.String("id", strconv.FormatInt(author.ID, 10)),
	)

	//localeCode := user.LanguageCode
	localeCode := localization.Russian
	locale := localization.Get(localeCode)

	query := u.CallbackData()
	log.Debug("Callback handled", slog.String("query", query))

	groupID, groupTitle, week, weekday := callback.ExtractFromScheduleCallback(query)

	doc, err := ssau.GetScheduleDocument(groupID, week)
	if err != nil {
		_, err = h.SendTextMessage(author.ID, locale.Message.NoScheduleFound, nil)
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
		GetScheduleNavigationMarkup(localeCode, groupID, groupTitle, week, weekday, favouriteButton))
	if errors.Is(err, ErrNoChanges) {
		cb := telegram.NewCallback(u.CallbackQuery.ID, locale.Message.NoChanges)
		_, err = h.bot.Request(cb)
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

	//localeCode := user.LanguageCode
	localeCode := localization.Russian
	locale := localization.Get(localeCode)

	query := u.CallbackData()
	log.Debug("Callback handled", slog.String("query", query))

	groupID, groupTitle := callback.ExtractFromFavouriteGroupCallback(query)

	err := h.storage.UpdateUserLinkedGroup(author.ID, groupID, groupTitle)
	if err != nil {
		log.Error("Failed to update user's group", sl.Err(err))
		cb := telegram.NewCallback(u.CallbackQuery.ID, locale.Message.Error)
		_, err = h.bot.Request(cb)
		if err != nil {
			log.Error("Failed to send callback", sl.Err(err))
		}
	}

	cb := telegram.NewCallback(u.CallbackQuery.ID, locale.Message.Success)
	_, err = h.bot.Request(cb)
	if err != nil {
		log.Error("Failed to send callback", sl.Err(err))
	}
}

func (h *Handler) OnCallbackAddCalendar(u telegram.Update) {
	author := u.CallbackQuery.From

	log := h.log.With(
		slog.String("op", "handler.OnCallbackAddCalendar"),
		slog.String("username", author.UserName),
		slog.String("id", strconv.FormatInt(author.ID, 10)),
	)

	//localeCode := user.LanguageCode
	localeCode := localization.Russian
	locale := localization.Get(localeCode)

	query := u.CallbackData()
	log.Debug("Callback handled", slog.String("query", query))

	groupID, languageCode := callback.ExtractFromAddCalendarCallback(query)

	calendar, err := h.storage.GetTrackedCalendar(groupID, languageCode)
	if err != nil {
		log.Error("Failed to get calendar", sl.Err(err))
	}

	var calendarID string
	if calendar == nil {
		createdCalendar, err := h.storage.CreateTrackedCalendar(groupID, languageCode)
		if err != nil {
			log.Error("Failed to create tracked calendar", sl.Err(err))
			cb := telegram.NewCallback(u.CallbackQuery.ID, locale.Message.Error)
			_, err = h.bot.Request(cb)
			if err != nil {
				log.Error("Failed to send callback", sl.Err(err))
			}
		}
		calendarID = createdCalendar.ID
	} else {
		calendarID = calendar.ID
	}

	log.Debug("Created tracked calendar", slog.String("id", calendarID))

	_, err = icalendar.WriteNextNWeeksScheduleToFile(calendarID, groupID, languageCode, 4)
	if err != nil {
		log.Error("Failed to create file with calendar", sl.Err(err))
		cb := telegram.NewCallback(u.CallbackQuery.ID, locale.Message.Error)
		_, err = h.bot.Request(cb)
		if err != nil {
			log.Error("Failed to send callback", sl.Err(err))
		}
		return
	}

	file, err := os.Open("./calendars/" + calendarID + ".ics")
	if err != nil {
		log.Error("Failed to open calendar file", sl.Err(err))
		cb := telegram.NewCallback(u.CallbackQuery.ID, locale.Message.Error)
		_, err = h.bot.Request(cb)
		if err != nil {
			log.Error("Failed to send callback", sl.Err(err))
		}
		return
	}
	defer file.Close()

	fileReader := telegram.FileReader{
		Name:   "calendar.ics",
		Reader: file,
	}

	document := telegram.NewDocument(author.ID, fileReader)
	document.Caption = locale.Message.YourCalendar

	_, err = h.bot.Send(document)
	if err != nil {
		log.Error("Failed to send message with document", sl.Err(err))
		cb := telegram.NewCallback(u.CallbackQuery.ID, locale.Message.Error)
		_, err = h.bot.Request(cb)
		if err != nil {
			log.Error("Failed to send callback", sl.Err(err))
		}
		return
	}
}

func (h *Handler) OnCallbackSetLanguage(u telegram.Update) {
	author := u.CallbackQuery.From

	log := h.log.With(
		slog.String("op", "handler.OnCallbackSetLanguage"),
		slog.String("username", author.UserName),
		slog.String("id", strconv.FormatInt(author.ID, 10)),
	)

	//localeCode := user.LanguageCode
	localeCode := localization.Russian
	locale := localization.Get(localeCode)

	query := u.CallbackData()
	log.Debug("Callback handled", slog.String("query", query))

	languageCode := callback.ExtractFromSetLanguageCallback(query)

	err := h.storage.UpdateUserLanguage(author.ID, languageCode)
	if err != nil {
		cb := telegram.NewCallback(u.CallbackQuery.ID, locale.Message.Error)
		_, err = h.bot.Request(cb)
		if err != nil {
			log.Error("Failed to send callback", sl.Err(err))
		}
		return
	}

	cb := telegram.NewCallback(u.CallbackQuery.ID, locale.Message.Success)
	_, err = h.bot.Request(cb)
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

	//localeCode := user.LanguageCode
	localeCode := localization.Russian
	locale := localization.Get(localeCode)

	query := u.CallbackData()
	log.Debug("Callback handled", slog.String("query", query))

	groupID, groupTitle := callback.ExtractFromScheduleTodayCallback(query)
	weekday := ssau.GetWeekday(0)

	doc, err := ssau.GetScheduleDocument(groupID, 0)
	if err != nil {
		_, err = h.SendTextMessage(author.ID, locale.Message.NoScheduleFound, nil)
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
		GetScheduleNavigationMarkup(localeCode, groupID, groupTitle, week, weekday, favouriteButton))
	if errors.Is(err, ErrNoChanges) {
		cb := telegram.NewCallback(u.CallbackQuery.ID, locale.Message.NoChanges)
		_, err = h.bot.Request(cb)
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

	//localeCode := user.LanguageCode
	localeCode := localization.Russian
	locale := localization.Get(localeCode)

	query := u.CallbackData()
	log.Debug("Callback handled", slog.String("query", query))

	err := h.storage.UpdateUserStage(author.ID, postgres.StageNone)
	if err != nil {
		log.Error("Failed to update user's stage", sl.Err(err))
		cb := telegram.NewCallback(u.CallbackQuery.ID, locale.Message.FailedToCancel)
		_, err = h.bot.Request(cb)
		if err != nil {
			log.Error("Failed to send callback", sl.Err(err))
		}
	}

	cb := telegram.NewCallback(u.CallbackQuery.ID, locale.Message.Cancelled)
	_, err = h.bot.Request(cb)
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

	//localeCode := user.LanguageCode
	localeCode := localization.Russian
	locale := localization.Get(localeCode)

	query := u.CallbackData()
	log.Debug("Callback handled", slog.String("query", query))

	content, exists := h.storage.GetAnnouncementMessage(author.ID)
	if !exists {
		log.Error("Content for announcement doesn't exists")
		cb := telegram.NewCallback(u.CallbackQuery.ID, locale.Message.AnnouncementMessageIsEmpty)
		_, err := h.bot.Request(cb)
		if err != nil {
			log.Error("Failed to send callback", sl.Err(err))
		}
	}

	users, err := h.storage.GetAllUsers()
	if err != nil {
		log.Error("Failed to get all users", sl.Err(err))
		cb := telegram.NewCallback(u.CallbackQuery.ID, locale.Message.CantStartAnnouncement)
		_, err = h.bot.Request(cb)
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

	_, err = h.EditMessageText(u.CallbackQuery.Message, locale.Message.AnnouncementCompleted, nil)
	if err != nil {
		log.Error("Failed to send message", sl.Err(err))
	}
}

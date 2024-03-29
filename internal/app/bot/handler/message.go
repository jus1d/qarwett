package handler

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"qarwett/internal/app/localization"
	"qarwett/internal/app/schedule"
	"qarwett/internal/app/ssau"
	"qarwett/internal/storage/postgres"
	"qarwett/pkg/logger/sl"
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

	//localeCode := author.LanguageCode
	localeCode := localization.Russian
	locale := localization.Get(localeCode)

	user, err := h.storage.GetUserByTelegramID(author.ID)
	if err != nil {
		log.Error("Failed to get user from database", sl.Err(err))
		_, err = h.SendTextMessage(author.ID, locale.Message.UseRestart, nil)
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
		}
		return
	}

	//localeCode = user.LanguageCode
	localeCode = localization.Russian
	locale = localization.Get(localeCode)

	if user.Stage == postgres.StageWaitingAnnouncementMessage {
		content := u.Message.Text
		_, err = h.SendTextMessage(author.ID, locale.Message.AnnouncementCheck(content), GetMarkupCheckAnnouncement(localeCode))
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
			return
		}
		h.storage.SetAnnouncementMessage(author.ID, content)
		return
	}

	query := u.Message.Text

	groups, err := ssau.GetGroupBySearchQuery(query)
	if len(groups) == 0 || err != nil {
		_, err = h.SendTextMessage(author.ID, locale.Message.NoGroupFound, nil)
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
		}
		return
	}

	if len(groups) > 1 {
		markup := GetMarkupFromGroupList(groups)
		_, err = h.SendTextMessage(author.ID, locale.Message.ChooseGroup, markup)
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
		}
		return
	}

	group := groups[0]
	doc, err := ssau.GetScheduleDocument(group.ID, 0)
	if err != nil {
		_, err = h.SendTextMessage(author.ID, locale.Message.NoScheduleFound, nil)
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
			return
		}
	}
	timetable, week := ssau.Parse(doc)

	user, err = h.storage.GetUserByTelegramID(author.ID)
	if err != nil {
		log.Error("Failed to get user from storage", sl.Err(err))
	}

	favouriteButton := err == nil && user.LinkedGroupID != group.ID

	weekday := ssau.GetWeekday(0)
	content := schedule.ParseScheduleToMessageTextWithHTML(group.ID, group.Title, schedule.Day{
		Date:  timetable.StartDate.AddDate(0, 0, weekday),
		Pairs: timetable.Pairs[weekday],
	})

	_, err = h.SendTextMessage(author.ID, content, GetScheduleNavigationMarkup(localeCode, group.ID, group.Title, week, weekday, favouriteButton))
	if err != nil {
		log.Error("Failed to send message", sl.Err(err))
		return
	}
}

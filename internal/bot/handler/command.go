package handler

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"qarwett/internal/lib/logger/sl"
	"qarwett/internal/locale"
	"qarwett/internal/schedule"
	"qarwett/internal/ssau"
	"qarwett/internal/storage/postgres"
	"strconv"
)

func (h *Handler) OnCommandStart(u telegram.Update) {
	author := u.Message.From

	log := h.log.With(
		slog.String("op", "handler.OnCommandStart"),
		slog.String("username", author.UserName),
		slog.String("id", strconv.FormatInt(author.ID, 10)),
	)

	log.Debug("Command triggered: /start")

	userExists := h.storage.IsUserExists(author.ID)
	if !userExists {
		id, err := h.storage.CreateUser(author.ID, author.UserName, author.FirstName, author.LastName, author.LanguageCode)
		if err != nil {
			log.Error("Failed to save user", sl.Err(err))
		} else {
			log.Debug("User saved", slog.String("id", id))
		}
	}

	message := telegram.NewMessage(u.Message.Chat.ID, locale.PhraseGreeting(locale.RU))

	message.ParseMode = telegram.ModeHTML

	_, err := h.bot.Send(message)
	if err != nil {
		log.Error("Failed to send greeting message", sl.Err(err))
	}
}

func (h *Handler) OnCommandAdmin(u telegram.Update) {
	author := u.Message.From

	log := h.log.With(
		slog.String("op", "handler.OnCommandAdmin"),
		slog.String("username", author.UserName),
		slog.String("id", strconv.FormatInt(author.ID, 10)),
	)

	log.Debug("Command triggered: /a")

	user, err := h.storage.GetUserByTelegramID(author.ID)
	if err != nil {
		log.Error("Failed to get user from database", sl.Err(err))
		return
	}

	if !user.IsAdmin {
		log.Info("Admin command triggerred by not admin", slog.String("command", u.Message.Text))
		return
	}

	_, err = h.SendTextMessage(author.ID, locale.PhraseAdminCommands(locale.RU), nil)
	if err != nil {
		log.Error("Failed to send message", sl.Err(err))
	}
}

func (h *Handler) OnCommandAnnounce(u telegram.Update) {
	author := u.Message.From

	log := h.log.With(
		slog.String("op", "handler.OnCommandAnnounce"),
		slog.String("username", author.UserName),
		slog.String("id", strconv.FormatInt(author.ID, 10)),
	)

	log.Debug("Command triggered: /announce")

	user, err := h.storage.GetUserByTelegramID(author.ID)
	if err != nil {
		log.Error("Failed to get user from database", sl.Err(err))
		return
	}

	if !user.IsAdmin {
		log.Info("Admin command triggered by not admin", slog.String("command", u.Message.Text))
		return
	}

	err = h.storage.UpdateUserStage(author.ID, postgres.StageWaitingAnnouncementMessage)
	if err != nil {
		_, err = h.SendTextMessage(author.ID, locale.PhraseCantStartAnnouncement(locale.RU), nil)
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
		}
		log.Error("Failed to update user's stage", sl.Err(err))
		return
	}

	_, err = h.SendTextMessage(author.ID, locale.PhraseAnnouncementRequest(locale.RU), GetMarkupCancel(locale.RU))
	if err != nil {
		log.Error("Failed to send message", sl.Err(err))
		return
	}
}

func (h *Handler) OnCommandUsers(u telegram.Update) {
	author := u.Message.From

	log := h.log.With(
		slog.String("op", "handler.OnCommandUsers"),
		slog.String("username", author.UserName),
		slog.String("id", strconv.FormatInt(author.ID, 10)),
	)

	log.Debug("Command triggered: /users")

	user, err := h.storage.GetUserByTelegramID(author.ID)
	if err != nil {
		log.Error("Failed to get user from database", sl.Err(err))
		return
	}

	if !user.IsAdmin {
		log.Info("Admin command triggered by not admin", slog.String("command", u.Message.Text))
		return
	}

	users, err := h.storage.GetAllUsers()
	if err != nil {
		log.Error("Failed to get users from database", sl.Err(err))
		return
	}

	_, err = h.SendTextMessage(author.ID, locale.PhraseUsersCommand(locale.RU, len(users)), nil)
	if err != nil {
		log.Error("Failed to send message", sl.Err(err))
	}
}

func (h *Handler) OnCommandToday(u telegram.Update) {
	author := u.Message.From

	log := h.log.With(
		slog.String("op", "handler.OnCommandToday"),
		slog.String("username", author.UserName),
		slog.String("id", strconv.FormatInt(author.ID, 10)),
	)

	log.Debug("Command triggered: /today")

	user, err := h.storage.GetUserByTelegramID(author.ID)
	if err != nil {
		log.Error("Failed to get user from database")
		_, err = h.SendTextMessage(author.ID, locale.PhraseUseRestart(locale.RU), nil)
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
			return
		}
	}

	if user.LinkedGroupID == 0 {
		_, err = h.SendTextMessage(author.ID, locale.PhraseCantFoundYourGroup(locale.RU), nil)
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
		}
		return
	}

	groupID := user.LinkedGroupID
	doc, err := ssau.GetScheduleDocument(groupID, 0)
	if err != nil {
		_, err = h.SendTextMessage(author.ID, locale.PhraseNoScheduleFound(locale.RU), nil)
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
			return
		}
	}
	timetable, week := ssau.Parse(doc)

	weekday := ssau.GetWeekday(0)
	content := schedule.ParseScheduleToMessageTextWithHTML(user.LinkedGroupTitle, schedule.Day{
		Date:  timetable.StartDate.AddDate(0, 0, weekday),
		Pairs: timetable.Pairs[weekday],
	})

	// TODO: Somehow get group title here
	_, err = h.SendTextMessage(author.ID, content, GetScheduleNavigationMarkup(groupID, user.LinkedGroupTitle, week, weekday, false))
	if err != nil {
		log.Error("Failed to send message", sl.Err(err))
		return
	}
}

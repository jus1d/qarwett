package handler

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"qarwett/internal/lib/logger/sl"
	"qarwett/internal/locale"
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

	message := telegram.NewMessage(u.Message.Chat.ID, locale.GetPhraseGreeting(locale.RU))

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
		log.Debug("Admin command triggerred by not admin")
		return
	}

	_, err = h.SendTextMessage(author.ID, locale.GetPhraseAdminCommands(locale.RU), nil)
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

	err := h.storage.UpdateUserStage(author.ID, postgres.StageWaitingAnnouncementMessage)
	if err != nil {
		_, err = h.SendTextMessage(author.ID, locale.GetPhraseCantStartAnnouncement(locale.RU), nil)
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
		}
		log.Error("Failed to update user's stage", sl.Err(err))
		return
	}

	_, err = h.SendTextMessage(author.ID, locale.GetPhraseAnnouncementRequest(locale.RU), GetMarkupCancel(locale.RU))
	if err != nil {
		log.Error("Failed to send message", sl.Err(err))
		return
	}
}

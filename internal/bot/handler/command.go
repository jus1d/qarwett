package handler

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"qarwett/internal/lib/locale"
	"qarwett/internal/lib/logger/sl"
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

	message := telegram.NewMessage(u.Message.Chat.ID, locale.GetPhraseGreeting("ru"))

	message.ParseMode = telegram.ModeHTML

	_, err := h.bot.Send(message)
	if err != nil {
		log.Error("Failed to send greeting message", sl.Err(err))
	}
}

package handler

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"qarwett/internal/lib/logger/sl"
	"strconv"
)

func (h *Handler) HandleStart(u telegram.Update) {
	author := u.Message.From

	log := h.log.With(
		slog.String("op", "handler.HandleStart"),
		slog.String("username", author.UserName),
		slog.String("id", strconv.FormatInt(author.ID, 10)),
	)

	log.Debug("Command triggered: /start")

	message := telegram.NewMessage(u.Message.Chat.ID,
		"<b>Hello, here you can view your timetable</b>\n\n"+
			"Just type your group 👇")

	message.ParseMode = telegram.ModeHTML

	_, err := h.bot.Send(message)
	if err != nil {
		log.Error("Failed to send greeting message", sl.Err(err))
	}
}

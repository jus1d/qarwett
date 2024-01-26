package handler

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"qarwett/internal/storage/postgres"
)

type Handler struct {
	log     *slog.Logger
	bot     *telegram.BotAPI
	storage *postgres.Storage
}

func New(log *slog.Logger, bot *telegram.BotAPI, storage *postgres.Storage) *Handler {
	return &Handler{
		log:     log,
		bot:     bot,
		storage: storage,
	}
}

func (h *Handler) SendTextMessage(chatID int64, content string, markup interface{}) (telegram.Message, error) {
	msg := telegram.NewMessage(chatID, content)
	msg.ParseMode = telegram.ModeHTML
	msg.ReplyMarkup = markup

	return h.bot.Send(msg)
}

func (h *Handler) EditMessageText(message *telegram.Message, content string, markup telegram.InlineKeyboardMarkup) (telegram.Message, error) {
	c := telegram.NewEditMessageText(message.Chat.ID, message.MessageID, content)
	c.ReplyMarkup = &markup
	c.ParseMode = telegram.ModeHTML

	return h.bot.Send(c)
}

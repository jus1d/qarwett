package handler

import (
	"errors"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"qarwett/internal/storage/postgres"
	"strings"
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

func (h *Handler) EditMessageText(message *telegram.Message, content string, markup *telegram.InlineKeyboardMarkup) (telegram.Message, error) {
	if message.Text == strings.TrimRight(RemoveHTML(content), "\n") {
		return *message, ErrNoChanges
	}
	c := telegram.NewEditMessageText(message.Chat.ID, message.MessageID, content)
	c.ReplyMarkup = markup
	c.ParseMode = telegram.ModeHTML

	msg, err := h.bot.Send(c)
	if err != nil && strings.Contains(err.Error(), "message is not modified") {
		return *message, ErrNoChanges
	}
	return msg, err
}

func RemoveHTML(s string) string {
	flag := true
	var builder strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == '<' {
			flag = false
		} else if s[i] == '>' {
			flag = true
		} else if flag {
			builder.WriteByte(s[i])
		}
	}
	return builder.String()
}

var (
	ErrNoChanges = errors.New("no changes")
)

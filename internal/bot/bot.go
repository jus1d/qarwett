package bot

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"qarwett/internal/bot/handler"
	"qarwett/internal/storage/postgres"
	"strings"
)

type Bot struct {
	env     string
	client  *telegram.BotAPI
	handler *handler.Handler
	log     *slog.Logger
}

func New(token string, env string, log *slog.Logger, storage *postgres.Storage) (*Bot, error) {
	bot, err := telegram.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	h := handler.New(log, bot, storage)

	return &Bot{
		env:     env,
		client:  bot,
		handler: h,
		log:     log,
	}, nil
}

func (b *Bot) Run() {
	log := b.log.With(slog.String("op", "bot.Run"), slog.String("env", b.env))

	updates := b.getUpdates()
	go b.handleUpdates(updates)

	log.Info("Bot successfully started", slog.String("username", b.client.Self.UserName))
}

func (b *Bot) getUpdates() telegram.UpdatesChannel {
	updates := telegram.NewUpdate(0)
	updates.Timeout = 30

	return b.client.GetUpdatesChan(updates)
}

func (b *Bot) handleUpdates(updates telegram.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil {
			switch update.Message.Command() {
			case "start":
				b.handler.OnCommandStart(update)
			default:
				b.handler.OnNewMessage(update)
			}
		}
		if update.CallbackQuery != nil {
			data := update.CallbackData()
			if strings.HasPrefix(data, "schedule:") {
				b.handler.OnCallbackSchedule(update)
			}
		}
	}
}

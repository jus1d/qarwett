package bot

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"qarwett/internal/bot/handler"
	"qarwett/internal/storage/postgres"
	"strings"
)

// Bot struct implements a structure to call telegram API, use storage etc.
type Bot struct {
	env     string
	client  *telegram.BotAPI
	handler *handler.Handler
	log     *slog.Logger
}

// New returns a new *Bot instance. If Some services don't start, function will return an error.
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

// Run create a thread and run function to handle different updates, caught via telegram API.
func (b *Bot) Run() {
	log := b.log.With(slog.String("op", "bot.Run"), slog.String("env", b.env))

	updates := b.getUpdates()
	go b.handleUpdates(updates)

	log.Info("Bot successfully started", slog.String("username", b.client.Self.UserName))
}

// getUpdates returns channel with updates, caught via telegram API.
func (b *Bot) getUpdates() telegram.UpdatesChannel {
	updates := telegram.NewUpdate(0)
	updates.Timeout = 30

	return b.client.GetUpdatesChan(updates)
}

// handleUpdates is infinitely grab new updates from channel, and handle them.
func (b *Bot) handleUpdates(updates telegram.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil {
			switch update.Message.Command() {
			case "start":
				b.handler.OnCommandStart(update)
			case "a":
				b.handler.OnCommandAdmin(update)
			case "announce":
				b.handler.OnCommandAnnounce(update)
			case "users":
				b.handler.OnCommandUsers(update)
			default:
				b.handler.OnNewMessage(update)
			}
		}

		if update.CallbackQuery != nil {
			data := update.CallbackData()
			if strings.HasPrefix(data, "schedule-daily:") {
				b.handler.OnCallbackSchedule(update)
			} else if strings.HasPrefix(data, "schedule-today:") {
				b.handler.OnCallbackScheduleToday(update)
			} else if strings.HasPrefix(data, "favourite-group") {
				b.handler.OnCallbackFavouriteGroup(update)
			} else if data == "cancel" {
				b.handler.OnCallbackCancel(update)
			} else if data == "approve-announcement" {
				b.handler.OnCallbackAnnouncementApprove(update)
			}
		}
	}
}

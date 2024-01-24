package bot

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"os"
	"os/signal"
	"qarwett/internal/bot/handler"
	"syscall"
)

type Bot struct {
	client  *telegram.BotAPI
	handler *handler.Handler
	log     *slog.Logger
}

func New(token string, log *slog.Logger) (*Bot, error) {
	bot, err := telegram.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	h := handler.New(log, bot)

	return &Bot{
		client:  bot,
		handler: h,
		log:     log,
	}, nil
}

func (b *Bot) Run() {
	log := b.log.With(slog.String("op", "bot.Run"))

	u := telegram.NewUpdate(0)
	u.Timeout = 60

	go b.handleUpdates()

	log.Info("Bot successfully started", slog.String("username", b.client.Self.UserName))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}

func (b *Bot) getUpdates() telegram.UpdatesChannel {
	updatesConfig := telegram.NewUpdate(0)
	updatesConfig.Timeout = 30

	return b.client.GetUpdatesChan(updatesConfig)
}

func (b *Bot) handleUpdates() {
	updates := b.getUpdates()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "start":
			b.handler.HandleStart(update)
		default:
			b.handler.HandleMessage(update)
		}
	}
}

func (b *Bot) SendTextMessage(chatID int64, content string, markup interface{}) (telegram.Message, error) {
	msg := telegram.NewMessage(chatID, content)
	msg.ParseMode = telegram.ModeHTML
	msg.ReplyMarkup = markup

	return b.client.Send(msg)
}

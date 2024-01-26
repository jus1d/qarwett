package handler

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"qarwett/internal/lib/logger/sl"
	"qarwett/internal/schedule"
	"qarwett/internal/ssau"
	"strconv"
)

type Handler struct {
	log *slog.Logger
	bot *telegram.BotAPI
}

func New(log *slog.Logger, bot *telegram.BotAPI) *Handler {
	return &Handler{
		log: log,
		bot: bot,
	}
}

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

func (h *Handler) HandleMessage(u telegram.Update) {
	author := u.Message.From

	log := h.log.With(
		slog.String("op", "handler.HandleMessage"),
		slog.String("username", author.UserName),
		slog.String("id", strconv.FormatInt(author.ID, 10)),
	)

	log.Debug("Message handled", slog.String("content", u.Message.Text))

	query := u.Message.Text

	groups, err := ssau.GetGroupBySearchQuery(query)
	if len(groups) == 0 || err != nil {
		_, err := h.SendTextMessage(author.ID, "Can't found group '"+query+"'.", nil)
		if err != nil {
			log.Error("Failed to send message")
		}
		return
	}

	// TODO(#3): Create button with different groups, to get user an ability to choose
	group := groups[0]
	doc, err := ssau.GetScheduleDocument(group.ID, 26)
	if err != nil {
		_, err = h.SendTextMessage(author.ID, "Can't get a schedule. Sorry!", nil)
		if err != nil {
			log.Error("Failed to send message")
			return
		}
	}
	timetable := ssau.Parse(doc)
	weekday := ssau.GetWeekday(0)

	content := schedule.ParseScheduleToMessageTextWithHTML(timetable[weekday])

	_, err = h.SendTextMessage(author.ID, content, nil)
	if err != nil {
		log.Error("Failed to send message")
		return
	}
}

func (h *Handler) SendTextMessage(chatID int64, content string, markup interface{}) (telegram.Message, error) {
	msg := telegram.NewMessage(chatID, content)
	msg.ParseMode = telegram.ModeHTML
	msg.ReplyMarkup = markup

	return h.bot.Send(msg)
}

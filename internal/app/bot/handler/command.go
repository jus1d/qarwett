package handler

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"qarwett/internal/app/localization"
	"qarwett/internal/app/schedule"
	"qarwett/internal/app/ssau"
	"qarwett/internal/storage/postgres"
	"qarwett/pkg/git"
	"qarwett/pkg/logger/sl"
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

	user, err := h.storage.GetUserByTelegramID(author.ID)
	if err != nil {
		user, err = h.storage.CreateUser(author.ID, author.UserName, author.FirstName, author.LastName, author.LanguageCode)
		if err != nil {
			log.Error("Failed to save user", sl.Err(err))
		} else {
			log.Debug("User saved", slog.String("id", user.ID))
		}
	}

	//localeCode := user.LanguageCode
	localeCode := localization.Russian
	locale := localization.Get(localeCode)

	message := telegram.NewMessage(u.Message.Chat.ID, locale.Message.Greeting)

	message.ParseMode = telegram.ModeHTML

	_, err = h.bot.Send(message)
	if err != nil {
		log.Error("Failed to send greeting message", sl.Err(err))
	}
}

func (h *Handler) OnCommandAbout(u telegram.Update) {
	author := u.Message.From

	log := h.log.With(
		slog.String("op", "handler.OnCommandAbout"),
		slog.String("username", author.UserName),
		slog.String("id", strconv.FormatInt(author.ID, 10)),
	)

	log.Debug("Command triggered: /about")

	commit := git.GetLatestCommit()

	//localeCode := user.LanguageCode
	localeCode := localization.Russian
	locale := localization.Get(localeCode)

	_, err := h.SendTextMessage(author.ID, locale.Message.About(commit), nil)
	if err != nil {
		log.Error("Failed to send message", sl.Err(err))
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

	//localeCode := user.LanguageCode
	localeCode := localization.Russian
	locale := localization.Get(localeCode)

	_, err = h.SendTextMessage(author.ID, locale.Message.AdminCommands, nil)
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

	//localeCode := user.LanguageCode
	localeCode := localization.Russian
	locale := localization.Get(localeCode)

	err = h.storage.UpdateUserStage(author.ID, postgres.StageWaitingAnnouncementMessage)
	if err != nil {
		_, err = h.SendTextMessage(author.ID, locale.Message.CantStartAnnouncement, nil)
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
		}
		log.Error("Failed to update user's stage", sl.Err(err))
		return
	}

	_, err = h.SendTextMessage(author.ID, locale.Message.RequestAnnouncement, GetMarkupCancel(localeCode))
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

	//localeCode := user.LanguageCode
	localeCode := localization.Russian
	locale := localization.Get(localeCode)

	if !user.IsAdmin {
		log.Info("Admin command triggered by not admin", slog.String("command", u.Message.Text))
		return
	}

	users, err := h.storage.GetAllUsers()
	if err != nil {
		log.Error("Failed to get users from database", sl.Err(err))
		return
	}

	_, err = h.SendTextMessage(author.ID, locale.Message.UsersAmount(len(users)), nil)
	if err != nil {
		log.Error("Failed to send message", sl.Err(err))
	}
}

func (h *Handler) OnCommandLanguage(u telegram.Update) {
	author := u.Message.From

	log := h.log.With(
		slog.String("op", "handler.OnCommandLanguage"),
		slog.String("username", author.UserName),
		slog.String("id", strconv.FormatInt(author.ID, 10)),
	)

	log.Debug("Command triggered: /language")

	//localeCode := author.LanguageCode
	localeCode := localization.Russian
	locale := localization.Get(localeCode)

	_, err := h.storage.GetUserByTelegramID(author.ID)
	if err != nil {
		log.Error("Failed to get user from database")
		_, err = h.SendTextMessage(author.ID, locale.Message.UseRestart, nil)
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
		}
		return
	}

	//localeCode := user.LanguageCode
	localeCode = localization.Russian
	locale = localization.Get(localeCode)

	_, err = h.SendTextMessage(author.ID, locale.Message.ChooseLanguage, GetLanguagesMarkup())
	if err != nil {
		log.Error("Failed to send message", sl.Err(err))
		return
	}
}

func (h *Handler) OnCommandSoon(u telegram.Update) {
	author := u.Message.From

	log := h.log.With(
		slog.String("op", "handler.OnCommandSoon"),
		slog.String("username", author.UserName),
		slog.String("id", strconv.FormatInt(author.ID, 10)),
	)

	log.Debug("Command triggered: /soon")

	//localeCode := author.LanguageCode
	localeCode := localization.Russian
	locale := localization.Get(localeCode)

	user, err := h.storage.GetUserByTelegramID(author.ID)
	if err != nil {
		log.Error("Failed to get user from database")
		_, err = h.SendTextMessage(author.ID, locale.Message.UseRestart, nil)
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
		}
		return
	}

	//localeCode := user.LanguageCode
	localeCode = localization.Russian
	locale = localization.Get(localeCode)

	if user.LinkedGroupID == 0 {
		log.Error("Failed to get user from database")
		_, err = h.SendTextMessage(author.ID, locale.Message.NoLinkedGroup, nil)
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
		}
		return
	}

	groupID := user.LinkedGroupID
	doc, err := ssau.GetScheduleDocument(groupID, 0)
	if err != nil {
		_, err = h.SendTextMessage(author.ID, locale.Message.NoScheduleFound, nil)
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
			return
		}
	}
	timetable, week := ssau.Parse(doc)
	weekday := ssau.GetWeekday(0)
	for weekday < 7 {
		if len(timetable.Pairs[weekday]) != 0 {
			break
		}

		weekday++

		if weekday == 7 {
			doc, err = ssau.GetScheduleDocument(groupID, week+1)
			if err != nil {
				_, err = h.SendTextMessage(author.ID, locale.Message.NoScheduleFound, nil)
				if err != nil {
					log.Error("Failed to send message", sl.Err(err))
					return
				}
			}
			timetable, week = ssau.Parse(doc)
			weekday = 0
		}
	}

	content := schedule.ParseScheduleToMessageTextWithHTML(user.LinkedGroupID, user.LinkedGroupTitle, schedule.Day{
		Date:  timetable.StartDate.AddDate(0, 0, weekday),
		Pairs: timetable.Pairs[weekday],
	})

	_, err = h.SendTextMessage(author.ID, content, GetScheduleNavigationMarkup(localeCode, groupID, user.LinkedGroupTitle, week, weekday, false))
	if err != nil {
		log.Error("Failed to send message", sl.Err(err))
		return
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

	//localeCode := author.LanguageCode
	localeCode := localization.Russian
	locale := localization.Get(localeCode)

	user, err := h.storage.GetUserByTelegramID(author.ID)
	if err != nil {
		log.Error("Failed to get user from database")
		_, err = h.SendTextMessage(author.ID, locale.Message.UseRestart, nil)
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
		}
		return
	}

	//localeCode := user.LanguageCode
	localeCode = localization.Russian
	locale = localization.Get(localeCode)

	if user.LinkedGroupID == 0 {
		_, err = h.SendTextMessage(author.ID, locale.Message.CantFoundYourGroup, nil)
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
		}
		return
	}

	groupID := user.LinkedGroupID
	doc, err := ssau.GetScheduleDocument(groupID, 0)
	if err != nil {
		_, err = h.SendTextMessage(author.ID, locale.Message.NoScheduleFound, nil)
		if err != nil {
			log.Error("Failed to send message", sl.Err(err))
			return
		}
	}
	timetable, week := ssau.Parse(doc)

	weekday := ssau.GetWeekday(0)
	content := schedule.ParseScheduleToMessageTextWithHTML(user.LinkedGroupID, user.LinkedGroupTitle, schedule.Day{
		Date:  timetable.StartDate.AddDate(0, 0, weekday),
		Pairs: timetable.Pairs[weekday],
	})

	_, err = h.SendTextMessage(author.ID, content, GetScheduleNavigationMarkup(localeCode, groupID, user.LinkedGroupTitle, week, weekday, false))
	if err != nil {
		log.Error("Failed to send message", sl.Err(err))
		return
	}
}

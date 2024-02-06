package app

import (
	"log/slog"
	"os"
	"os/signal"
	telegram "qarwett/internal/app/bot"
	"qarwett/internal/config"
	"qarwett/internal/pkg/app/updater"
	"qarwett/internal/storage/postgres"
	"qarwett/pkg/logger"
	"qarwett/pkg/logger/sl"
	"syscall"
)

type App struct {
	config *config.Config
	log    *slog.Logger
}

func New() *App {
	cfg := config.MustLoad()
	log := logger.Init(cfg.Env)

	return &App{
		config: cfg,
		log:    log,
	}
}

func (a *App) Run() {
	storage, err := postgres.New(a.config.Postgres)
	if err != nil {
		a.log.Error("Failed to connect to database", sl.Err(err))
		return
	}

	bot, err := telegram.New(a.config.Telegram.Token, a.config.Env, a.log, storage)
	if err != nil {
		a.log.Error("Can't create a bot instance", sl.Err(err))
		return
	}

	go bot.Run()

	iCalendarUpdater := updater.New(a.config, a.log, storage)
	go iCalendarUpdater.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	a.log.Info("Bot is shutting down")

	err = storage.Close()
	if err != nil {
		a.log.Error("Failed to close postgresql", sl.Err(err))
	}

	a.log.Info("Postgres connection closed")
}

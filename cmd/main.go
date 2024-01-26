package main

import (
	"os"
	"os/signal"
	"qarwett/internal/bot"
	"qarwett/internal/config"
	"qarwett/internal/lib/logger"
	"qarwett/internal/lib/logger/sl"
	"qarwett/internal/storage/postgres"
	"syscall"
)

// TODO(#2): Add image generating for weekly timetable

func main() {
	cfg := config.MustLoad()

	log := logger.Init(cfg.Env)

	storage, err := postgres.New(cfg.Postgres)
	if err != nil {
		log.Error("Failed to connect to database", sl.Err(err))
		return
	}

	b, err := bot.New(cfg.Telegram.Token, log, storage)
	if err != nil {
		log.Error("Can't create a bot instance", sl.Err(err))
		return
	}

	go b.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Info("Bot is shutting down")

	err = storage.Close()
	if err != nil {
		log.Error("Failed to close postgresql", sl.Err(err))
	}

	log.Info("Postgres connection closed")
}

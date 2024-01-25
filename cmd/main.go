package main

import (
	"qarwett/internal/bot"
	"qarwett/internal/config"
	"qarwett/internal/lib/logger"
	"qarwett/internal/lib/logger/sl"
)

// TODO(#2): Add image generating for weekly timetable

func main() {
	cfg := config.MustLoad()

	log := logger.Init(cfg.Env)

	b, err := bot.New(cfg.Telegram.Token, log)
	if err != nil {
		log.Error("Can't create a bot instance", sl.Err(err))
		return
	}

	b.Run()
}

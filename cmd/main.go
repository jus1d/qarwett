package main

import (
	"qarwett/internal/bot"
	"qarwett/internal/config"
	"qarwett/internal/lib/logger"
	"qarwett/internal/lib/logger/sl"
)

func main() {
	cfg := config.MustLoad()

	log := logger.Init(cfg.Env)

	if err != nil {
		log.Error("Can't create a bot instance", sl.Err(err))
		return
	}

	b.Run()
}

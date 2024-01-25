package main

import (
	"fmt"
	"qarwett/internal/bot"
	"qarwett/internal/config"
	"qarwett/internal/lib/logger"
	"qarwett/internal/lib/logger/sl"
	schedule2 "qarwett/internal/schedule"
	"qarwett/internal/ssau"
)

func main2() {
	cfg := config.MustLoad()

	log := logger.Init(cfg.Env)

	b, err := bot.New(cfg.Telegram.Token, log)
	if err != nil {
		log.Error("Can't create a bot instance", sl.Err(err))
		return
	}

	b.Run()
}

func main() {
	groups, _ := ssau.GetGroupBySearchQuery("4102-030302D")
	schedule, _ := ssau.Parse(groups[0].ID, 12)

	for i := 0; i < len(schedule[schedule2.Thursday]); i++ {
		fmt.Println(schedule[schedule2.Thursday][i])
	}
}

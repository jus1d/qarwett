package logger

import (
	"log/slog"
	"os"
	"qarwett/internal/config"
	"qarwett/internal/lib/logger/prettyslog"
)

func Init(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case config.EnvLocal:
		log = InitPretty()
	case config.EnvDevelopment:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case config.EnvProduction:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}

func InitPretty() *slog.Logger {
	opts := prettyslog.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	prettyHandler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(prettyHandler)
}

package logger

import (
	"log/slog"
	"os"
	"qarwett/internal/config"
	"qarwett/pkg/logger/prettyslog"
)

// Init initialize a *slog.Logger instance for logging, without pretty formatting for production and development builds.
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

// InitPretty initialize a *slog.Logger instance for logging, with pretty formatting for local builds.
func InitPretty() *slog.Logger {
	opts := prettyslog.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	prettyHandler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(prettyHandler)
}

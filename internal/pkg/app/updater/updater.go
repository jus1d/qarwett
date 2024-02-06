package updater

import (
	"github.com/robfig/cron/v3"
	"log/slog"
	"qarwett/internal/app/icalendar"
	"qarwett/internal/config"
	"qarwett/internal/storage/postgres"
	"qarwett/pkg/logger/sl"
)

const (
	Schedule           = "0 12,0 * * *" // Every day at 12 AM and PM
	AmountWeeksToTrack = 4
)

type Updater struct {
	config  *config.Config
	log     *slog.Logger
	storage *postgres.Storage
}

func New(cfg *config.Config, log *slog.Logger, stg *postgres.Storage) *Updater {
	return &Updater{
		config:  cfg,
		log:     log,
		storage: stg,
	}
}

func (u *Updater) Run() {
	log := u.log.With(slog.String("service", "app.Updater"))

	c := cron.New()

	_, err := c.AddFunc(Schedule, func() {
		calendars, err := u.storage.GetAllTrackedCalendars()
		if err != nil {
			log.Error("Can't get all tracked calendars", sl.Err(err))
		}

		for _, calendar := range calendars {
			_, _ = icalendar.WriteNextNWeeksScheduleToFile(calendar.ID, calendar.GroupID, calendar.LanguageCode, AmountWeeksToTrack)
		}
	})
	if err != nil {
		log.Error("Can't add CRON worker", sl.Err(err))
	}

	c.Start()
}

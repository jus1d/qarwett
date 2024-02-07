package updater

import (
	"github.com/robfig/cron/v3"
	"log/slog"
	"qarwett/internal/app/icalendar"
	"qarwett/internal/config"
	"qarwett/internal/storage/postgres"
	"qarwett/pkg/logger/sl"
)

const Schedule = "0 12,0 * * *" // Every day at 12 AM and PM

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

	u.getAndUpdateCalendars()

	c := cron.New()

	_, err := c.AddFunc(Schedule, func() {
		u.getAndUpdateCalendars()
	})
	if err != nil {
		log.Error("Can't add CRON worker", sl.Err(err))
	}

	c.Start()
}

func (u *Updater) getAndUpdateCalendars() {
	log := u.log.With(slog.String("service", "app.Updater"))

	calendars, err := u.storage.GetAllTrackedCalendars()
	if err != nil {
		log.Error("Can't get all tracked calendars", sl.Err(err))
	}

	for _, calendar := range calendars {
		file, err := icalendar.WriteNextNWeeksScheduleToFile(calendar.ID, calendar.GroupID,
			calendar.LanguageCode, u.config.ICalendar.Updater.WeeksToTrack)
		if err != nil {
			log.Error("Failed to update calendar file", sl.Err(err))
		} else {
			log.Info("Calendar file updated", slog.String("file", file))
		}
	}
}

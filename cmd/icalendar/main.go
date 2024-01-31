package main

import (
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"os/signal"
	"qarwett/internal/app/icalendar"
	"qarwett/internal/app/locale"
	"syscall"
)

const Schedule = "0 12,0 * * *" // Every day at 12 AM and PM

func main() {
	log.Printf("Schedule Format: '%s'\n", Schedule)
	groupIDs := []int64{755922237}

	c := cron.New()
	for _, groupID := range groupIDs {
		_, err := c.AddFunc(Schedule, func() {
			updateICalendarFile(groupID)
		})
		if err != nil {
			log.Printf("ERROR: %s\n", err.Error())
		} else {
			log.Printf("Schedule requested for group with ID: %d\n", groupID)
		}
	}

	c.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}

func updateICalendarFile(groupID int64) {
	_, _ = icalendar.WriteNextNWeeksScheduleToFile(groupID, locale.RU, 4)
}

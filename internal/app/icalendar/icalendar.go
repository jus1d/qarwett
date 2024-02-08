package icalendar

import (
	"fmt"
	"log"
	"os"
	"qarwett/internal/app/locale"
	"qarwett/internal/app/schedule"
	"qarwett/internal/app/ssau"
	"time"
)

const CalendarsDir = "calendars"

var pairPositionToMinutesFromDayStart = map[int]int{
	0: 480,
	1: 585,
	2: 690,
	3: 810,
	4: 915,
	5: 1020,
	6: 1125,
	7: 1225,
}

func WriteNextNWeeksScheduleToFile(filename string, groupID int64, languageCode string, n int) (string, error) {
	if _, err := os.Stat(CalendarsDir); os.IsNotExist(err) {
		_ = os.Mkdir(CalendarsDir, 0755)
	}
	filename = fmt.Sprintf("%s/%s.ics", CalendarsDir, filename)
	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content string
	addICalendarHeader(&content, languageCode)
	week := -1
	for i := 0; i < n; i++ {
		doc, err := ssau.GetScheduleDocument(groupID, week+1)
		if err != nil {
			log.Fatal(err)
		}
		var sch schedule.WeekPairs
		sch, week = ssau.Parse(doc)
		addICalendarSchedule(&content, sch, languageCode)
	}
	addICalendarFooter(&content)

	_, err = file.WriteString(content)
	if err != nil {
		return "", err
	}

	err = file.Sync()
	if err != nil {
		return "", err
	}

	if err = file.Close(); err != nil {
		return "", err
	}

	return filename, nil
}

func addICalendarSchedule(content *string, schedule schedule.WeekPairs, languageCode string) {
	for i := 0; i < len(schedule.Pairs); i++ {
		day := schedule.Pairs[i]
		dayStart := schedule.StartDate.AddDate(0, 0, i)
		for j := 0; j < len(day); j++ {
			pair := day[j]
			start := dayStart.Add(time.Duration(pairPositionToMinutesFromDayStart[pair.Position]) * time.Minute)
			addICalendarEvent(content, pair, start, languageCode)
		}
	}
}

func addICalendarEvent(content *string, pair schedule.Pair, start time.Time, languageCode string) {
	languageCode = locale.RU

	end := start.Add(95 * time.Minute)
	*content += fmt.Sprintf("\nBEGIN:VEVENT\n")
	*content += fmt.Sprintf("DTSTART:%s\n", start.UTC().Format("20060102T150405"))
	*content += fmt.Sprintf("DTEND:%s\n", end.UTC().Format("20060102T150405"))
	*content += fmt.Sprintf("DESCRIPTION:%s", pair.Staff.Name)
	if pair.Subgroup != 0 {
		*content += fmt.Sprintf(" %s: %d", locale.ScheduleSubgroup(languageCode), pair.Subgroup)
	}
	*content += "\n"
	*content += fmt.Sprintf("LOCATION:%s %s %s\n", schedule.FullPairTypes[pair.Type], locale.ScheduleIn(languageCode), pair.Place)
	*content += fmt.Sprintf("SUMMARY:%s", pair.Title)
	if pair.Subgroup != 0 {
		*content += fmt.Sprintf(" (%d)", pair.Subgroup)
	}
	*content += "\n"
	*content += "END:VEVENT\n"
}

func addICalendarHeader(content *string, languageCode string) {
	calendarName := locale.ScheduleCalendarName(locale.RU)
	*content += fmt.Sprintf("BEGIN:VCALENDAR\nVERSION:2.0\nCALSCALE:GREGORIAN\nMETHOD:PUBLISH\nX-WR-CALNAME:%s\nX-WR-TIMEZONE:Europe/Samara\n", calendarName)
}

func addICalendarFooter(content *string) {
	*content += "\nEND:VCALENDAR"
}

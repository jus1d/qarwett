package icalendar

import (
	"fmt"
	"os"
	"qarwett/internal/schedule"
	"time"
)

const CalendarsDir = ".calendars"

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

func WriteScheduleToFile(groupID int64, schedule schedule.WeekPairs) (string, error) {
	if _, err := os.Stat(CalendarsDir); os.IsNotExist(err) {
		_ = os.Mkdir(CalendarsDir, 0755)
	}
	filename := fmt.Sprintf("%s/%d.ics", CalendarsDir, groupID)
	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content string
	addICalendarHeader(&content)
	addICalendarSchedule(&content, schedule)
	addICalendarFooter(&content)

	_, err = file.WriteString(content)
	if err != nil {
		return "", err
	}
	return filename, nil
}

func addICalendarSchedule(content *string, schedule schedule.WeekPairs) {
	for i := 0; i < len(schedule.Pairs); i++ {
		day := schedule.Pairs[i]
		dayStart := schedule.StartDate.AddDate(0, 0, i)
		for j := 0; j < len(day); j++ {
			pair := day[j]
			start := dayStart.Add(time.Duration(pairPositionToMinutesFromDayStart[pair.Position]) * time.Minute)
			addICalendarEvent(content, pair, start)
		}
	}
}

func addICalendarEvent(content *string, pair schedule.Pair, start time.Time) {
	end := start.Add(95 * time.Minute)
	*content += fmt.Sprintf("BEGIN:VEVENT\n")
	*content += fmt.Sprintf("DTSTART:%s\n", start.UTC().Format("20060102T150405"))
	*content += fmt.Sprintf("DTEND:%s\n", end.UTC().Format("20060102T150405"))
	*content += fmt.Sprintf("DESCRIPTION:%s", pair.Staff.Name)
	if pair.Subgroup != 0 {
		*content += fmt.Sprintf(" Подгруппа: %d", pair.Subgroup)
	}
	*content += "\n"
	*content += fmt.Sprintf("LOCATION:%s в %s\n", schedule.FullPairTypes[pair.Type], pair.Place)
	*content += fmt.Sprintf("SUMMARY:%s", pair.Title)
	if pair.Subgroup != 0 {
		*content += fmt.Sprintf(" (%d)", pair.Subgroup)
	}
	*content += "\n"
	*content += "END:VEVENT\n"
}

func addICalendarHeader(content *string) {
	calendarName := "University Schedule"
	*content += fmt.Sprintf("BEGIN:VCALENDAR\nVERSION:2.0\nCALSCALE:GREGORIAN\nMETHOD:PUBLISH\nX-WR-CALNAME:%s\nX-WR-TIMEZONE:Europe/Samara\n", calendarName)
}

func addICalendarFooter(content *string) {
	*content += "END:VCALENDAR"
}

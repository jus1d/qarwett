package icalendar

import (
	"fmt"
	"log"
	"os"
	"qarwett/internal/schedule"
	"time"
)

var PairPositionToMinutesFromDayStart = map[int]int{
	0: 480,
	1: 585,
	2: 690,
	3: 810,
	4: 915,
	5: 1020,
	6: 1125,
	7: 1225,
}

func ScheduleToFile(groupID int64, schedule schedule.WeekPairs) {
	if _, err := os.Stat(".calendars"); os.IsNotExist(err) {
		_ = os.Mkdir(".calendars", 0755)
	}
	filename := fmt.Sprintf(".calendars/%d.icalendar", groupID)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writeSchedule(file, schedule)
}

func writeSchedule(file *os.File, schedule schedule.WeekPairs) {
	writeHeaderICS(file)

	for i := 0; i < len(schedule.Pairs); i++ {
		day := schedule.Pairs[i]
		date := schedule.StartDate.AddDate(0, 0, i)
		for j := 0; j < len(day); j++ {
			pair := day[j]
			start := date.Add(time.Duration(PairPositionToMinutesFromDayStart[pair.Position]) * time.Minute)
			writeEvent(file, pair, start)
		}
	}

	writeFooterICS(file)
}

func writeEvent(file *os.File, pair schedule.Pair, startTime time.Time) {
	endTime := startTime.UTC().Add(95 * time.Minute)
	content := fmt.Sprintf("BEGIN:VEVENT\n")
	content += fmt.Sprintf("DTSTART:%s\n", startTime.UTC().Format("20060102T150405"))
	content += fmt.Sprintf("DTEND:%s\n", endTime.Format("20060102T150405"))
	content += fmt.Sprintf("DESCRIPTION:%s", pair.Staff.Name)
	if pair.Subgroup != 0 {
		content += fmt.Sprintf(" Подгруппа: %d", pair.Subgroup)
	}
	content += "\n"
	writeToFile(file, content)
	content = fmt.Sprintf("LOCATION:%s в %s\n", schedule.FullPairTypes[pair.Type], pair.Place)
	content += fmt.Sprintf("SUMMARY:%s", pair.Title)
	if pair.Subgroup != 0 {
		content += fmt.Sprintf(" (%d)", pair.Subgroup)
	}
	content += "\n"
	content += "END:VEVENT\n"

	writeToFile(file, content)
}

func writeHeaderICS(file *os.File) {
	calendarName := "University Schedule"
	content := fmt.Sprintf("BEGIN:VCALENDAR\nVERSION:2.0\nCALSCALE:GREGORIAN\nMETHOD:PUBLISH\nX-WR-CALNAME:%s\nX-WR-TIMEZONE:Europe/Samara\n", calendarName)
	writeToFile(file, content)
}

func writeFooterICS(file *os.File) {
	writeToFile(file, "END:VCALENDAR")
}

func writeToFile(file *os.File, content string) {
	_, err := file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

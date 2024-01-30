package main

import (
	"fmt"
	"log"
	"qarwett/internal/icalendar"
	"qarwett/internal/locale"
	"qarwett/internal/ssau"
)

func main() {
	query := "6101-020302D"

	groups, err := ssau.GetGroupBySearchQuery(query)
	if err != nil {
		log.Fatal(err)
	}
	group := groups[0]

	filename, err := icalendar.WriteNextNWeeksScheduleToFile(group.ID, locale.RU, 4)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Calendar created at: %s\n", filename)
}

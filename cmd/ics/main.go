package main

import (
	"log"
	"qarwett/internal/icalendar"
	"qarwett/internal/ssau"
)

func main() {
	query := "6101-020302D"

	groups, err := ssau.GetGroupBySearchQuery(query)
	if err != nil {
		log.Fatal(err)
	}
	group := groups[0]

	doc, err := ssau.GetScheduleDocument(group.ID, 26)
	if err != nil {
		log.Fatal(err)
	}

	schedule, _ := ssau.Parse(doc)

	icalendar.ScheduleToFile(group.ID, schedule)
}

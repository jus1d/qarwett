package ssau

import (
	"github.com/PuerkitoBio/goquery"
	"qarwett/internal/schedule"
	"strings"
)

var PairColors = map[string]schedule.PairType{
	"lesson-color-type-1": schedule.Lecture,
	"lesson-color-type-2": schedule.Lab,
	"lesson-color-type-3": schedule.Practice,
	"lesson-color-type-4": schedule.Other,
	"lesson-color-type-5": schedule.Exam,
	"lesson-color-type-6": schedule.Consult,
	"lesson-color-type-8": schedule.Exam,
}

func Parse(groupId int64, week int) ([][]schedule.Pair, error) {
	doc, err := GetScheduleDocument(groupId, week)
	if err != nil {
		return nil, err
	}

	pairs := make([][]schedule.Pair, 6)
	for i := 0; i < len(pairs); i++ {
		pairs[i] = make([]schedule.Pair, 0)
	}

	doc.Find(".schedule__item:not(.schedule__head)").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "" {
			return
		}

		dayOfWeek := i % 6
		ordinalNumber := i / 6

		s.Find(".schedule__lesson").Each(func(j int, s *goquery.Selection) {
			pair := ParsePair(s, ordinalNumber)
			pairs[dayOfWeek] = append(pairs[dayOfWeek], pair)
		})
	})

	return pairs, nil
}

func ParsePair(doc *goquery.Selection, ordinalNumber int) schedule.Pair {
	discipline := doc.Find(".schedule__discipline")
	title := discipline.Text()
	classAttributes := strings.Split(discipline.AttrOr("class", "lesson-color-type-4"), " ")
	pairColor := classAttributes[len(classAttributes)-1]
	pairType := PairColors[pairColor]

	place := doc.Find(".schedule__place").Text()

	teacher := doc.Find(".schedule__teacher").Find("a.caption-text").Text()

	var groups []string
	doc.Find(".schedule__groups").Find("a.schedule__group").Each(func(i int, s *goquery.Selection) {
		groups = append(groups, s.Text())
	})

	return schedule.Pair{
		OrdinalNumber: ordinalNumber,
		Type:          pairType,
		Title:         title,
		Place:         place,
		Staff:         teacher,
		Groups:        groups,
		SubGroup:      0,
	}
}

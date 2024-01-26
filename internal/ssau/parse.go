package ssau

import (
	"github.com/PuerkitoBio/goquery"
	"qarwett/internal/schedule"
	"strconv"
	"strings"
)

var PairColors = map[string]schedule.PairType{
	"1": schedule.Lecture,
	"2": schedule.Lab,
	"3": schedule.Practice,
	"4": schedule.Other,
	"5": schedule.Exam,
	"6": schedule.Consultation,
	"8": schedule.Test,
}

func Parse(doc *goquery.Document) [][]schedule.Pair {
	pairs := make([][]schedule.Pair, 6)
	for i := 0; i < len(pairs); i++ {
		pairs[i] = make([]schedule.Pair, 0)
	}

	doc.Find(".schedule__item:not(.schedule__head)").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "" {
			return
		}

		weekday := i % 6
		pos := i / 6

		s.Find(".schedule__lesson").Each(func(j int, s *goquery.Selection) {
			pair := parsePair(s, pos)
			pairs[weekday] = append(pairs[weekday], pair)
		})
	})

	return pairs
}

func parsePair(doc *goquery.Selection, pos int) schedule.Pair {
	discipline := doc.Find(".schedule__discipline")
	title := discipline.Text()
	classAttributes := strings.Split(discipline.AttrOr("class", "4"), " ")
	pairColorClass := classAttributes[len(classAttributes)-1]
	parts := strings.Split(pairColorClass, "-")
	pairColor := parts[len(parts)-1]
	pairType := PairColors[pairColor]

	place := doc.Find(".schedule__place").Text()

	teacherName := doc.Find(".schedule__teacher").Find("a.caption-text").Text()
	teacherURL := doc.Find(".schedule__teacher").Find("a.caption-text").AttrOr("href", "https://ssau.ru")

	var groups []schedule.Group
	doc.Find(".schedule__groups").Find("a.schedule__group").Each(func(i int, s *goquery.Selection) {
		groupUrl := s.AttrOr("href", "https://ssau.ru")
		groups = append(groups, schedule.Group{
			ID:    GetIdFromURL(groupUrl),
			Title: strings.TrimSpace(s.Text()),
		})
	})

	subgroupText := doc.Find("div.schedule__groups span.caption-text").Text()

	var subgroup int
	if subgroupText == "" {
		subgroup = 0
	} else {
		parts := strings.Split(strings.Replace(subgroupText, " ", "", -1), ":")
		subgroup, _ = strconv.Atoi(parts[len(parts)-1])
	}

	return schedule.Pair{
		Position: pos,
		Type:     pairType,
		Title:    strings.TrimSpace(title),
		Place:    strings.TrimSpace(place),
		Staff: schedule.Staff{
			ID:   GetIdFromURL(teacherURL),
			Name: strings.TrimSpace(teacherName),
		},
		Groups:   groups,
		SubGroup: subgroup,
	}
}

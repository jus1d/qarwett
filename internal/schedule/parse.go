package schedule

import (
	"fmt"
	"qarwett/internal/locale"
)

// TODO(#10): Add group title to schedule representation

// ParseScheduleToMessageTextWithHTML parses a daily schedule, to text message for telegram.
func ParseScheduleToMessageTextWithHTML(groupTitle string, schedule Day) string {
	pairs := schedule.Pairs
	date := schedule.Date
	months := []string{"", "января", "февраля", "марта", "апреля", "мая", "июня", "июля", "августа", "сентября", "октября", "сентября", "декабря"}
	if len(pairs) == 0 {
		return locale.PhraseForFreeDay("ru", date.Day(), int(date.Month()))
	}

	content := fmt.Sprintf("Расписание на <b>%d %s</b>\n", date.Day(), months[date.Month()])

	if groupTitle != "" {
		content += fmt.Sprintf("<b>Группа:</b> %s\n", groupTitle)
	}
	content += "\n"

	for i := 0; i < len(pairs); i++ {
		cur := pairs[i]

		if i == 0 || cur.Position != pairs[i-1].Position {
			content += fmt.Sprintf("<b>%s</b>\n", Timetable[cur.Position])
		}
		if i == 0 || cur.Position != pairs[i-1].Position || cur.Title != pairs[i-1].Title {
			content += fmt.Sprintf("<b>%s:</b> %s\n", FullPairTypes[cur.Type], cur.Title)
		}

		if cur.Subgroup != 0 {
			content += fmt.Sprintf("Подгруппа: %d\n", cur.Subgroup)
		}
		content += fmt.Sprintf("%s\n", cur.Place)
		if cur.Staff.Name != "" {
			content += fmt.Sprintf("<a href=\"https://ssau.ru/rasp?staffId=%d\">%s</a>\n", cur.Staff.ID, cur.Staff.Name)
		}
		if i < len(pairs)-1 && cur.Position == pairs[i+1].Position {
			content += "|\n"
		} else if i != len(pairs)-1 {
			content += "\n"
		}
	}

	return content
}

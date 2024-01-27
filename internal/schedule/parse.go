package schedule

import (
	"fmt"
)

func ParseScheduleToMessageTextWithHTML(schedule Day) string {
	// TODO(#7): Add date and weekday
	pairs := schedule.Pairs
	date := schedule.Date
	months := []string{"", "января", "февраля", "марта", "апреля", "мая", "июня", "июля", "августа", "сентября", "октября", "сентября", "декабря"}
	if len(pairs) == 0 {
		return fmt.Sprintf("<b>%d %s - свободный день</b>", date.Day(), months[date.Month()])
	}

	content := fmt.Sprintf("<b>Расписание на %d %s</b>\n\n", date.Day(), months[date.Month()])

	for i := 0; i < len(pairs); i++ {
		cur := pairs[i]

		if i == 0 || cur.Position != pairs[i-1].Position || cur.Title != pairs[i-1].Title {
			content += fmt.Sprintf("<b>%s</b>\n", Timetable[cur.Position])
			content += fmt.Sprintf("<b>%s:</b> %s\n", FullPairTypes[cur.Type], cur.Title)
		}

		if cur.SubGroup != 0 {
			content += fmt.Sprintf("Подгруппа: %d\n", cur.SubGroup)
		}
		content += fmt.Sprintf("%s\n", cur.Place)
		content += fmt.Sprintf("<a href=\"https://ssau.ru/rasp?staffId=%d\">%s</a>\n", cur.Staff.ID, cur.Staff.Name)
		if i < len(pairs)-1 && cur.Position == pairs[i+1].Position {
			content += "|\n"
		} else if i != len(pairs)-1 {
			content += "\n"
		}
	}

	return content
}

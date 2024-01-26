package schedule

import (
	"fmt"
)

func ParseScheduleToMessageTextWithHTML(schedule []Pair) string {
	if len(schedule) == 0 {
		return "<b>Расписание пустое</b>"
	}

	content := "<b>Расписание</b>\n\n"

	for i := 0; i < len(schedule); i++ {
		cur := schedule[i]

		content += fmt.Sprintf("<b>%s</b>\n", Timetable[cur.Position])
		content += fmt.Sprintf("<b>%s:</b> %s\n", FullPairTypes[cur.Type], cur.Title)
		content += fmt.Sprintf("%s\n", cur.Place)
		content += fmt.Sprintf("<a href=\"https://ssau.ru/rasp?staffId=%d\">%s</a>\n",
			cur.Staff.ID, cur.Staff.Name)
		if i != len(schedule)-1 {
			content += "\n"
		}
	}

	return content
}

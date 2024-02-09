package schedule

import (
	"fmt"
	"qarwett/internal/app/localization"
)

// ParseScheduleToMessageTextWithHTML parses a daily schedule, to text message for telegram.
func ParseScheduleToMessageTextWithHTML(groupID int64, groupTitle string, schedule Day) string {
	//localeCode := user.LanguageCode
	localeCode := localization.Russian
	locale := localization.Get(localeCode)

	pairs := schedule.Pairs
	date := schedule.Date
	var content string
	if groupTitle != "" {
		content += fmt.Sprintf("<b><u><a href=\"https://ssau.ru/rasp?groupId=%d\">%s:</a></u></b>\n\n", groupID, groupTitle)
	}

	months := []string{"", "января", "февраля", "марта", "апреля", "мая", "июня", "июля", "августа", "сентября", "октября", "сентября", "декабря"}
	if len(pairs) == 0 {
		return content + locale.Message.FreeDay(date.Day(), int(date.Month()))
	}

	content += fmt.Sprintf("Расписание на <b>%d %s</b>\n\n", date.Day(), months[date.Month()])

	for i := 0; i < len(pairs); i++ {
		cur := pairs[i]

		if i == 0 || cur.Position != pairs[i-1].Position {
			content += fmt.Sprintf("<b>%s</b>\n", Timetable[cur.Position])
		}
		if i == 0 || cur.Position != pairs[i-1].Position || cur.Title != pairs[i-1].Title {
			content += fmt.Sprintf("<b>%s</b>\n", cur.Title) // Another type
			//content += fmt.Sprintf("<b>%s:</b> %s\n", FullPairTypes[cur.Type], cur.Title)
		}

		if cur.Subgroup != 0 {
			content += fmt.Sprintf("Подгруппа: %d\n", cur.Subgroup)
		}
		//content += fmt.Sprintf("%s\n", cur.Place)
		content += fmt.Sprintf("%s в %s\n", FullPairTypes[cur.Type], cur.Place) // Another type

		if cur.Staff.Name != "" {
			content += fmt.Sprintf("<b><a href=\"https://ssau.ru/rasp?staffId=%d\">%s</a></b>\n", cur.Staff.ID, cur.Staff.Name)
		}
		if i < len(pairs)-1 && cur.Position == pairs[i+1].Position {
			content += "|\n"
		} else if i != len(pairs)-1 {
			content += "\n"
		}
	}

	return content
}

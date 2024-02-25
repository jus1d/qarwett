package localization

import (
	"fmt"
	"qarwett/pkg/random"
)

var EnglishLocalization = Locale{
	Message: Messages{
		Greeting: "<b>Hello, here you can take a quick look at your schedule <s>and go get some sleep</s></b>\n\n" +
			"👇Just type your group",
		RequestAnnouncement:        "Send me an announcement message",
		CantStartAnnouncement:      "An error occurred while creating an announcement!",
		AnnouncementCompleted:      "<b>Announcement completely sent to all users</b>",
		CantFoundYourGroup:         "☹️Can't found your group\n\nAdd your group <b>to favourites</b>, to use this command",
		NoGroupFound:               "☹️There are no groups at your request",
		ChooseGroup:                "🤔<b>Choose a group</b>",
		NoScheduleFound:            "🚨<b>Can't found schedule!</b>",
		NoLinkedGroup:              "☹️You have no <b>favourite</b> group",
		NoChanges:                  "No changes",
		Error:                      "Error",
		Success:                    "Success",
		Cancelled:                  "Cancelled",
		FailedToCancel:             "Failed to cancel",
		AnnouncementMessageIsEmpty: "Announcement message not found. Try again",
		UseRestart:                 "Restart needed → <b>/start</b>",
		AdminCommands: "<b>List of admin commands:</b>\n\n" +
			"<b>/announce</b> - Send an announcement message to all users\n" +
			"<b>/users</b> - View users counter",
		ChooseLanguage: "<b>Choose your language:</b>",
		YourCalendar:   "Your schedule file",
		languageCode:   English,
	},
	Button: Buttons{
		Today:       "Today",
		Favourite:   "To Favourites",
		Cancel:      "Cancel",
		Approve:     "Approve",
		AddCalendar: "Add to Calendar",
	},
	Schedule: Schedule{
		CalendarName: "University Schedule",
		Subgroup:     "Subgroup",
		In:           "in",
	},
	languageCode: English,
}

var phrasesForFreeDayEN = []string{
	"<b>%d %s</b> - a day of idleness",
	"<b>%d %s</b> even the alarm clock will be sleeping",
	"Plan for <b>%d %s</b> - just have no plans",
	"<b>%d %s</b> is a perfect day to head to Sochi",
	"<b>%d %s</b> - I allow you to be lazy",
	"<b>%d %s</b> - a day of exploration, investigate where the second socks constantly disappear",
	"Although it's <b>%d %s</b> off, stop lying on the couch",
}

func aboutEN(commit string) string {
	if commit == "" {
		commit = "xx"
	}

	return "<b>qarweTT</b> - take a quick look at schedule, and go back to sleep\n\n" +
		"With any questions, you can write me <b>@jus1d</b>\n\n" +
		"Current build version: <b>" + commit + "</b>"
}

func announcementCheckEN(content string) string {
	return "<b>Your announcement message:</b>\n\n" + content + "\n\n" +
		"Is everything correct?"
}

func usersAmountEN(amount int) string {
	return fmt.Sprintf("<b>Total users:</b> %d", amount)
}

func freeDayEN(day int, month int) string {
	monthsEN := []string{"", "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	return fmt.Sprintf(random.Choice(phrasesForFreeDayEN), day, monthsEN[month])
}

package localization

const (
	Russian = "ru"
	English = "en"
)

type Locale struct {
	Messages     messages
	Buttons      buttons
	languageCode string
}

type messages struct {
	Greeting                 string
	AnnouncementRequest      string
	CantStartAnnouncement    string
	AnnouncementCompleted    string
	CantFoundYourGroup       string
	NoGroupFound             string
	ChooseGroup              string
	NoScheduleFound          string
	NoChanges                string
	Error                    string
	Success                  string
	Cancelled                string
	FailedToCancel           string
	EmptyAnnouncementMessage string
	UseRestart               string
	AdminCommands            string
}

type buttons struct {
	Today       string
	Favourite   string
	Cancel      string
	Approve     string
	AddCalendar string
}

type Localization interface {
	About(commit string) string
	AnnouncementCheck(content string) string
	UsersAmount(amount int) string
	FreeDay(day int, month int) string
}

func Get(languageCode string) Locale {
	switch languageCode {
	case Russian:
		return RussianLocalization
	case English:
		return EnglishLocalization
	default:
		return EnglishLocalization
	}
}

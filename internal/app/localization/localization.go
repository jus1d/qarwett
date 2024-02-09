package localization

const (
	Russian = "ru"
	English = "en"
)

type Locale struct {
	Message      Messages
	Button       Buttons
	languageCode string
}

type Messages struct {
	Greeting                   string
	RequestAnnouncement        string
	CantStartAnnouncement      string
	AnnouncementCompleted      string
	CantFoundYourGroup         string
	NoGroupFound               string
	ChooseGroup                string
	NoScheduleFound            string
	NoChanges                  string
	Error                      string
	Success                    string
	Cancelled                  string
	FailedToCancel             string
	AnnouncementMessageIsEmpty string
	UseRestart                 string
	AdminCommands              string
	languageCode               string
}

type Buttons struct {
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

func (m *Messages) About(commit string) string {
	switch m.languageCode {
	case Russian:
		return aboutRU(commit)
	case English:
		return aboutEN(commit)
	default:
		return aboutEN(commit)
	}
}

func (m *Messages) AnnouncementCheck(content string) string {
	switch m.languageCode {
	case Russian:
		return announcementCheckRU(content)
	case English:
		return announcementCheckEN(content)
	default:
		return announcementCheckEN(content)
	}
}

func (m *Messages) UsersAmount(amount int) string {
	switch m.languageCode {
	case Russian:
		return usersAmountRU(amount)
	case English:
		return usersAmountEN(amount)
	default:
		return usersAmountEN(amount)
	}
}

func (m *Messages) FreeDay(day int, month int) string {
	switch m.languageCode {
	case Russian:
		return freeDayRU(day, month)
	case English:
		return freeDayEN(day, month)
	default:
		return freeDayEN(day, month)
	}
}

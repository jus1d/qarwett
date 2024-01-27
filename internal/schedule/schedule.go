package schedule

import "time"

type PairType string

const (
	Empty        PairType = "emp"
	Lecture      PairType = "lct"
	Practice     PairType = "prc"
	Lab          PairType = "lab"
	Other        PairType = "oth"
	Military     PairType = "mil"
	Exam         PairType = "exm"
	Consultation PairType = "cns"
	Test         PairType = "tst"
	CourseWork   PairType = "crs"
	Unknown      PairType = "unk"
)

const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

var Timetable = map[int]string{
	0: "8:00 - 9:35",
	1: "9:45 - 11:20",
	2: "11:30 - 13:05",
	3: "13:30 - 15:05",
	4: "15:15 - 16:50",
	5: "17:00 - 18:35",
	6: "18:45 - 20:15",
	7: "20:25 - 21:55",
}

var FullPairTypes = map[PairType]string{
	Empty:        "Пусто",
	Lecture:      "Лекция",
	Practice:     "Практика",
	Lab:          "Лабораторная",
	Other:        "Другое",
	Military:     "Военная кафедра",
	Exam:         "Экзамен",
	Consultation: "Консультация",
	Test:         "Зачет",
	CourseWork:   "Курсовая работа",
	Unknown:      "Неизвестно",
}

type Pair struct {
	Position int
	Type     PairType
	Title    string
	Place    string
	Staff    Staff
	Groups   []Group
	Subgroup int
}

type WeekDays struct {
	StartDate time.Time
	Days      []Day
}

type WeekPairs struct {
	StartDate time.Time
	Pairs     [][]Pair
}

type Day struct {
	Date  time.Time
	Pairs []Pair
}

type Staff struct {
	ID   int64
	Name string
}

type Group struct {
	ID    int64
	Title string
}

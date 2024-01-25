package schedule

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
}

// TODO: Add time for pairs 7+ to the timetable

type Pair struct {
	Position int
	Type     PairType
	Title    string
	Place    string
	Staff    Staff
	Groups   []Group
	SubGroup int
}

type Staff struct {
	ID   int64
	Name string
}

type Group struct {
	ID    int64
	Title string
}

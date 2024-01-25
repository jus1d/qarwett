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

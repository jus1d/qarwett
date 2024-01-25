package schedule

type PairType string

const (
	Lecture    PairType = "lct"
	Practice   PairType = "prc"
	Lab        PairType = "lab"
	Other      PairType = "oth"
	Military   PairType = "mil"
	Empty      PairType = "emp"
	Exam       PairType = "exm"
	Consult    PairType = "cns"
	CourseWork PairType = "crs"
	Test       PairType = "tst"
	Unknown    PairType = "unk"
)

type Pair struct {
	Type     PairType
	Title    string
	Place    string
	StaffID  int64
	GroupIDs []int64
}

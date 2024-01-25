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
	OrdinalNumber int
	Type          PairType
	Title         string
	Place         string
	Staff         string
	Groups        []string
	SubGroup      int
}

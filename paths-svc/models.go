package main

type objType int

const (
	city objType = iota
	militaryBase
	hospital
	AEC
	VEC
	TEC
)

func (o objType) String() string {
	return []string{"city", "militaryBase", "hospital", "AEC", "VEC", "TEC"}[o]
}

var objPriority = map[string]float64{
	"city":         4.5,
	"militaryBase": 1,
	"hospital":     3,
	"AEC":          19,
	"VEC":          7,
	"TEC":          9,
}

func Priority(o objType) float64 {
	return objPriority[o.String()]
}

type Path struct {
	From, To      Object
	Distance      float64 `json:"distance"`
	PriorityScore float64
}

type Object struct {
	Coords    Coordinates
	Paths     []Path
	TypeScore float64
}

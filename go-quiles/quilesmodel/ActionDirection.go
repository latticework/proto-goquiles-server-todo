package quilesmodel

type ActionDirection int

const (
	DirectionNone ActionDirection = iota
	To
	From
)

var actionDirections = [...]string{
	"DirectionNone",
	"To",
	"From",
}

func (d ActionDirection) String() string {
	return actionDirections[d]
}

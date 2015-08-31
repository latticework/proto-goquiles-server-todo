package quilesmodel

type ActionDirection int

const (
	None ActionDirection = iota
	To
	From
)

var actionDirections = [...]string{
	"None",
	"To",
	"From",
}

func (d ActionDirection) String() string {
	return actionDirections[d]
}

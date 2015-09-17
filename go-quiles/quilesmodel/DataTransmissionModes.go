package quilesmodel

type DataTransmissionModes int

const (
	ModeNone   DataTransmissionModes = 1 << iota
	ModeResult DataTransmissionModes = 1
	Notify     DataTransmissionModes = 1 << iota
	Patch
	Full
	All = ModeResult + Notify + Patch + Full
)

var dataTransmissionModes = map[DataTransmissionModes]string{
	ModeNone:   "ModeNone",
	ModeResult: "ModeResult",
	Notify:     "Notify",
	Patch:      "Patch",
	Full:       "Full",
	All:        "All",
}

func (v DataTransmissionModes) String() string {

	result := ""

	for index, f := range dataTransmissionModes {
		if v&f == f {
			name, ok := dataTransmissionModes[v]

			if ok {
				result += name
			}

			result += "(" + int(v) + ")"

			if index != 0 {
				result += ", "
			}
		}
	}

	return result
}

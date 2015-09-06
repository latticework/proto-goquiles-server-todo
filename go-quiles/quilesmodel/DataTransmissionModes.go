package quilesmodel

type DataTransmissionModes int

const (
	None   DataTransmissionModes = 1 << iota
	Result DataTransmissionModes = 1
	Notify DataTransmissionModes = 1 << iota
	Patch
	Full
	All = Result + Notify + Patch + Full
)

var dataTransmissionModes = map[DataTransmissionModes]string{
	None:   "None",
	Result: "Result",
	Notify: "Notify",
	Patch:  "Patch",
	Full:   "Full",
	All:    "All",
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

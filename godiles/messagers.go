package diles

type Messagers struct {
	slice []Messager
}

func NewMessager(slice ...Messager) Messagers {
	var messages []Message

	copy(messages, slice)

	return Messager{slice: messages}
}

func (m Messagers) Len() int { return len(m.slice) }
func (m Messagers) Less(i, j int) bool { return m.slice[i].Severity() < m.slice[j].Severity() }
func (m Messagers) Swap(i, j int) { m.slice[i], m.slice[j] = m.slice[j], m.slice[i] }

func []Messager Errors() {

}


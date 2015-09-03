package diles

import "sort"

type Messageators struct {
	slice []Messageator
}

func NewMessages(slice ...Messageator) *Messageators {
	var messages []Message

	copy(messages, slice)

	return &Messageator{slice: messages}
}

func (m *Messageators) Len() int      { return len(m.slice) }
func (m *Messageators) Swap(i, j int) { m.slice[i], m.slice[j] = m.slice[j], m.slice[i] }

type BySeverityDesc struct{ Messageators }

func (m BySeverityDesc) Less(i, j int) bool { return m.slice[i].Severity() < m.slice[j].Severity() }

func (m *Messageators) Errors() []Messageator {
	sorted := m.slice[:]
	sort.Stable(BySeverityDesc{sorted})

	for index, message := range sorted {
		if message.Severity() > Error {
			return sorted[:index]
		}
	}

	return sorted
}

func (m *Messageators) Messages() []Messageator {
	return m.slice[:]
}

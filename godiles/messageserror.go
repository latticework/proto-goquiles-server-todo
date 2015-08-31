package diles

import (
	"bytes"
	"sort"
)

type MessagesError struct {
	Messages Messagers
}

func (err *MessagesError) Error() string {
	var sorted Messagers
	copy(sorted, err.Messages)
	sort.Stable(BySeverityDesc(sorted))

	var messages bytes.Buffer
}

type BySeverityDesc struct{ Messagers }

func (m BySeverityDesc) Less(i, j int) bool {
	return m.Messagers[i].Severity() < m.Messagers[j].Severity()
}

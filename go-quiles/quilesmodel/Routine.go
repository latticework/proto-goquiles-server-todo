package quilesmodel

import "net/url"

type Routine struct {
	Url        url.URL
	Name       string
	EntryPoint url.URL
	Messages   map[string]RoutineMessage
}

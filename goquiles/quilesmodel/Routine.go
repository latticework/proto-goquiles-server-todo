package quilesmodel

import "net/url"

type Routine struct {
	Url     url.URL
	Name    string
	Actions map[string]RoutineAction
}

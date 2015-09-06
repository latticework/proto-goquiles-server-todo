package quilesmodel

type RestMethod struct {
	Method   string
	Routine  Routine
	Request  RestMethodRequest
	Response RestMethodResponse
}

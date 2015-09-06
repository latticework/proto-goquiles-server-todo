package svr

import (
	"github.com/latticework/proto-goquiles-server-todo/go-quiles"
	"log"
)

// Pepa is spanish for seed--as in chile seeds that make
// chile peppers so hot in chilequiles.
//
// Pepa is the core runtime--or kernel--for a quiles server
type Pepa struct {
}

type PepaConfig struct {
	Logger *log.Logger
}

func NewPepa(conf *PepaConfig) (*Pepa, error) {
	pepa := &Pepa{}

	return pepa, nil
}

func HandleMessage(message quiles.ServiceMessageator) {

}

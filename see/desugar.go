// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package see

import (
	"strconv"
)

func DesugarCircuit(cir *Circuit) *Circuit {
	var n int
	for _, m := range cir.Match {
		for j := 0; j < 2; j++ {
			switch t := m.Join[j].(type) {
			case *DesignJoin:
				name := "desugar" + strconv.Itoa(n)
				n++
				cir.Peer = append(cir.Peer, &Peer{ // create an anonymous peer for the design
					Name: name,
					Design: t.Design,
				})
				m.Join[j] = &PeerJoin{
					Peer: name,
					Valve: "",
				} // point the join to the anonymous peer
			}
		}
	}
	return cir
}

// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

type Orient map[Name]map[Name]struct{} // gate -> valve -> yes/no?

func (o Orient) Include(gate, valve Name) {
	valves, ok := o[gate]
	if !ok {
		valves = make(map[Name]struct{})
		o[gate] = valves
	}
	valves[valve] = struct{}{}
}

func (o Orient) Has(gate, valve Name) bool {
	if valves, ok := o[gate]; ok {
		if _, ok := valves[valve]; ok {
			return true
		}
	}
	return false
}

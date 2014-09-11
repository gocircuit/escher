// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

type Orient map[Name]map[Name]struct{} // image -> valve -> yes/no?

func (o Orient) Include(image, valve Name) {
	valves, ok := o[image]
	if !ok {
		valves = make(map[Name]struct{})
		o[image] = valves
	}
	valves[valve] = struct{}{}
}

func (o Orient) Has(image, valve Name) bool {
	if valves, ok := o[image]; ok {
		if _, ok := valves[valve]; ok {
			return true
		}
	}
	return false
}

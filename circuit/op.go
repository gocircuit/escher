// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

type X Circuit

// XXX: Copy and comparison conversions should traverse type renames!
func init() {
	x := X(New())
	var y interface{} = x
	_, ok := y.(Circuit)
	println("op.go:", ok)
}

// func (u *circuit) Match(image Name, meaning Meaning) bool {
// 	m, ok := u.image[image]
// 	if !ok {
// 		return false
// 	}
// 	return SameMeaning(m, meaning)
// }

// func (u Circuit) Sub(w Circuit) (include, exclude Circuit) {
// 	in, ex := u.circuit.Sub(w.circuit)
// 	return Circuit{in}, Circuit{ex}
// }

// func (u *circuit) Sub(w *circuit) (include, exclude *circuit) {
// 	include = newCircuit()
// 	for im1, x1 := range u.Images() {
// 		x2, ok := w[im1]
// 		if !ok {
// 			include.Include(im1, CopyMeaning(x1))
// 			continue
// 		}
// 		if SameMeaning(x1, x2) {
// 			continue
// 		}
// 		?? // recurse on circuits???
// 		exclude.Include(im1, CopyMeaning(x2))
// 		include.Include(im1, CopyMeaning(x1))
// 	}
// 	for im1, valves := range w.Reals() {
// 		for v1, re := range valves {
// 			im2, v2 := re.To()
// 			??
// 		}
// 	}
// 	return
// }

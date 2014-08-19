// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

func (nerve *EyeNerve) ReCognize(imp Impression) {
	<-nerve.connected
	ch := make(chan struct{})
	order := imp.Order()
	for _, f_ := range order {
		f := f_
		go func() {
			nerve.recognize[f.Valve()].ReCognize(f.Value())
			ch <- struct{}{}
		}()
	}
	for _ = range order {
		<-ch
	}
}

func (nerve *EyeNerve) cognizeWith(valve string, value interface{}) {
	<-nerve.connected
	nerve.Lock()
	nerve.age++
	nerve.memory.Show(nerve.age, valve, value)
	reply := nerve.formulate()
	nerve.Unlock()
	nerve.cognize(reply)
}

func (nerve *EyeNerve) formulate() Impression {
	var sorting = nerve.memory.Order()
	imp := MakeImpression()
	for i, f := range sorting {
		imp.Show(i, f.Valve(), f.Value())
	}
	return imp	
}

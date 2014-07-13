// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package record

// Record is a data structure modeled after:
//	http://research.microsoft.com/pubs/65409/scopedlabels.pdf
type Record map[string][]interface{}

// Make allocates a new record structure.
func Make() Record {
	return make(Record)
}

// Extends adds a new scope to the record with a given initial value.
func (rec Record) Extend(label string, value interface{}) Record {
	rec[label] = append(rec[label], value)
	return rec
}

// Restrict removes the label from the record.
func (rec Record) Restrict(label string) {
	scope := rec[label]
	if len(scope) == 1 {
		delete(rec, label)
		return
	}
	rec[label] = scope[:len(scope)-1]
}

// Copy copies just the high-level map of this record into a new one.
func (rec Record) Copy() Record {
	s := Make()
	for label, scope := range rec {
		s[label] = make([]interface{}, len(scope))
		copy(s[label], scope)
	}
	return s
}

// Flattens leaves only the top-level element of each scope in the record.
func (rec Record) Flatten() {
	for label, scope := range rec {
		rec[label] = []interface{}{scope[len(scope)-1]}
	}
}

func TranslateObservation(observation Record) (sense string, what interface{}, intelligible bool) {
	label, intelligible = observation["Sense"]
	if !intelligible { // no change in belief if input is unintelligible
		return "", nil, false
	}
	scope, intelligible = observation["What?"]
	if !intelligible { // no change in belief if input is unintelligible
		return "", nil, false
	}
	return label, scope[len(scope)-1], true
}

// Belief combined with an Observation produces a Theory.
func Generalize(belief, observation Record) Record {
	sense, what, intelligible := TranslateObservation(observation)
	if !intelligible { // no change in belief, if observation is unintelligible
		return belief
	}
	return belief.Copy().Extend(sense, what)
}

//  Theory combined with an Observation produces an Explanation.
func Explain(theory, observation Record) Record {
	sense, what, intelligible := TranslateObservation(observation)
	if !intelligible { // no change in belief, if observation is unintelligible
		return belief
	}
	for label, scope := range theory {
		if label == sense {
			if ?? {
			} else {
				// If prior theory does not address this category of observation, return confusion.
				return nil
			}
		}
	}
	// If observation was not found in theory, i.e. it is not explained by a consistent belief.
	// Return nil to signify confusion.
	return nil
}

//  Belief combined with Theory produces a Prediction.
func Predict(belief, theory Record) Record {
	diff := theory.Copy()
	for label, bs := range belief {
		ts, ok := diff[label]
		if !ok {
			xx
		}
	}
	??
	return diff
}

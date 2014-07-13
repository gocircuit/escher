// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package tree

// Branch is…
type Branch []interface{}

func (b Branch) Yield() interface{} {
	if len(b) == 0 {
		return nil
	}
	return b[len(b)-1]
}

// Tree is a data structure modeled after:
//	http://research.microsoft.com/pubs/65409/scopedlabels.pdf
type Tree map[string]Branch

// Make allocates a new tree structure.
func Make() Tree {
	return make(Tree)
}

// Grow adds a new scope to the tree with a given initial value.
func (tree Tree) Grow(name string, value interface{}) Tree {
	tree[name] = append(tree[name], value)
	return tree
}

// Restrict removes the name from the tree.
func (tree Tree) Restrict(name string) {
	scope := tree[name]
	if len(scope) == 1 {
		delete(tree, name)
		return
	}
	tree[name] = scope[:len(scope)-1]
}

// Copy copies just the high-level map of this tree into a new one.
func (tree Tree) Copy() Tree {
	s := Make()
	for name, scope := range tree {
		s[name] = make([]interface{}, len(scope))
		copy(s[name], scope)
	}
	return s
}

// Flattens leaves only the top-level element of each scope in the tree.
func (tree Tree) Flatten() {
	for name, scope := range tree {
		tree[name] = Branch{scope[len(scope)-1]}
	}
}

func TranslateObservation(observation Tree) (sense string, what interface{}, intelligible bool) {
	name, intelligible = observation["Sense"]
	if !intelligible { // no change in belief if input is unintelligible
		return "", nil, false
	}
	scope, intelligible = observation["What?"]
	if !intelligible { // no change in belief if input is unintelligible
		return "", nil, false
	}
	return name, scope[len(scope)-1], true
}

// Belief combined with an Observation produces a Theory.
func Generalize(belief, observation Tree) Tree {
	sense, what, intelligible := TranslateObservation(observation)
	if !intelligible { // no change in belief, if observation is unintelligible
		return belief
	}
	return belief.Copy().Grow(sense, what)
}

//  Theory combined with an Observation produces an Explanation.
func Explain(theory, observation Tree) Tree {
	sense, what, intelligible := TranslateObservation(observation)
	if !intelligible { // no change in belief, if observation is unintelligible
		return belief
	}
	for name, scope := range theory {
		if name == sense {
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
func Predict(belief, theory Tree) Tree {
	diff := theory.Copy()
	for name, bs := range belief {
		ts, ok := diff[name]
		if !ok {
			xx
		}
	}
	??
	return diff
}

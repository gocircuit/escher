// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package tree

import (
	"github.com/gocircuit/escher/think"
)

// Branch is…
type Branch []interface{}

func (b Branch) Yield() (interface{}, bool) {
	if len(b) == 0 {
		return nil, false
	}
	return b[len(b)-1], true
}

func SameYield(g, h Branch) bool {
	gy, gok := g.Yield()
	hy, hok := h.Yield()
	if gok != hok {
		return false
	}
	if !gok {
		return true
	}
	return think.Same(gy, hy)
}

// Tree is a data structure modeled after:
//	http://research.microsoft.com/pubs/65409/branchdlabels.pdf
type Tree map[string]Branch

// Make allocates a new tree structure.
func Make() Tree {
	return make(Tree)
}

// Grow adds a new branch to the tree with a given initial value.
func (tree Tree) Grow(name string, value interface{}) Tree {
	tree[name] = append(tree[name], value)
	return tree
}

// Forget removes the name from the tree.
func (tree Tree) Forget(name string) {
	branch := tree[name]
	if len(branch) == 1 {
		delete(tree, name)
		return
	}
	tree[name] = branch[:len(branch)-1]
}

// Copy copies just the high-level map of this tree into a new one.
func (tree Tree) Copy() Tree {
	s := Make()
	for name, branch := range tree {
		s[name] = make(Branch, len(branch))
		copy(s[name], branch)
	}
	return s
}

// Project leaves only the top-level element of each branch in the tree.
func (tree Tree) Project() (shadow Tree) {
	shadow = tree.Copy()
	for name, branch := range tree {
		shadow[name], _ = branch.Yield()
	}
	return
}

func (tree Tree) Mix(s Tree) (teach, learn Tree) { // (t-s, s-t) setwise
	teach, learn = Make(), Make()
	for name, branch := range tree {
		if idea, known := s[name]; !known || !SameYield(idea, branch) {
			teach[name], _ = branch.Yield()
		}
	}
	for name, idea := range s {
		if branch, know := tree[name]; !know || !SameYield(branch, idea) {
			learn[name], _ = idea.Yield()
		}
	}
	return
}

func TranslateObservation(observation Tree) (sense string, what interface{}, intelligible bool) {
	name, intelligible = observation["Sense"]
	if !intelligible { // no change in belief if input is unintelligible
		return "", nil, false
	}
	branch, intelligible = observation["What"]
	if !intelligible { // no change in belief if input is unintelligible
		return "", nil, false
	}
	what, _ = branch.Yield()
	return name, what, true
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
	for name, idea := range theory {
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
	teach, _ := theory.Mix(belief)
	switch len(teach) {
	case 0:
		??
	case 1:
		??
	default:
		??
	}
	panic(8)
}

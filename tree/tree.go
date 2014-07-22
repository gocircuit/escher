// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package tree

import (
	"reflect"
)

type Branch []Meaning

type Meaning interface{}

func (b Branch) Yield() (Meaning, bool) {
	if len(b) == 0 {
		return nil, false
	}
	return b[len(b)-1], true
}

func (b Branch) YieldNil() (m Meaning) {
	m, _ = b.Yield()
	return
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
	return Same(gy, hy)
}

func Same(v, w Meaning) bool {
	return reflect.DeepEqual(v, w)
}

// Tree is a data structure modeled after:
//	http://research.microsoft.com/pubs/65409/branchdlabels.pdf
type Tree map[Name]Branch

type Name interface{}

func Make() Tree {
	return make(Tree)
}

func Plant(name Name, value Meaning) Tree {
	return Make().Grow(name, value)
}

func (tree Tree) Grow(name Name, value Meaning) Tree {
	tree[name] = append(tree[name], value)
	return tree
}

func (tree Tree) At(name Name) Meaning {
	v, ok := tree[name]
	if !ok {
		panic(7)
	}
	return v
}

func (tree Tree) AtInt(name Name) int {
	v, ok := tree[name]
	if !ok {
		panic(7)
	}
	return v.YieldNil().(int)
}

func (tree Tree) AtString(name Name) string {
	v, ok := tree[name]
	if !ok {
		panic(7)
	}
	return v.YieldNil().(string)
}

func (tree Tree) AtName(name Name) Name {
	v, ok := tree[name]
	if !ok {
		panic(7)
	}
	return v.YieldNil().(Name)
}

// Forget removes the name from the tree.
func (tree Tree) Forget(name Name) {
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
	shadow = Make()
	for name, branch := range tree {
		shadow.Grow(name, branch.YieldNil())
	}
	return
}

func (tree Tree) Collapse() Tree {
	for name, branch := range tree {
		tree[name] = Branch{branch.YieldNil()}
	}
	return tree
}

func (tree Tree) Mix(s Tree) (teach, learn Tree) { // (t-s, s-t) setwise
	teach, learn = Make(), Make()
	for name, branch := range tree {
		if idea, known := s[name]; !known || !SameYield(idea, branch) {
			teach.Grow(name, branch.YieldNil())
		}
	}
	for name, idea := range s {
		if branch, know := tree[name]; !know || !SameYield(branch, idea) {
			learn.Grow(name, idea.YieldNil())
		}
	}
	return
}

// ConceptualizeObservation maps {x/y,…} to {Sense/x, What/y}
func ConceptualizeObservation(obs Tree) Tree {
	for sense, what := range obs {
		return Make().Grow("Sense", sense).Grow("What", what.YieldNil())
	}
	panic(8)
}

func DeConceptualizeObservation(observation Tree) Tree {
	name, intelligible := observation["Sense"]
	if !intelligible { // no change in belief if input is unintelligible
		return nil
	}
	branch, intelligible := observation["What"]
	if !intelligible { // no change in belief if input is unintelligible
		return nil
	}
	return Make().Grow(name.YieldNil(), branch.YieldNil())
}

// Belief combined with an Observation produces a Theory.
func Generalize(belief, observation Tree) (theory Tree) {
	obs := DeConceptualizeObservation(observation)
	for sense, what := range obs {
		return belief.Copy().Grow(sense, what)
	}
	panic(8)
}

//  Theory combined with an Observation produces a Belief.
func Explain(theory, observation Tree) (belief Tree) {
	belief, _ = theory.Mix(DeConceptualizeObservation(observation))
	return
}

//  Belief combined with Theory produces a Prediction.
func Predict(belief, theory Tree) (observation Tree) {
	teach, _ := theory.Mix(belief) // _ is the undiscovered misunderstanding in a conversation
	switch len(teach) {
	case 0:
		return nil
	case 1:
		return ConceptualizeObservation(teach)
	default:
		return nil // _ inaction in the face of overwheming difference between theory and belief
	}
	panic(8)
}

// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package draw

import (
	"fmt"
	"math"
	// "sync"
	"text/template"

	// "github.com/gocircuit/escher/faculty"
	// . "github.com/gocircuit/escher/image"
	// "github.com/gocircuit/escher/be"
	// "github.com/gocircuit/escher/plumb"
	"github.com/gocircuit/escher/understand"
)

type Circuit struct {
	Name string
	Peer []*Peer
	Match []*Match
}

type Vector struct {
	X, Y float64
}

type Peer struct {
	ID string
	Name, Design string
	Weight float64
	Anchor Vector
	Angle float64 // Angle of origin-center line in [0,2*Pi]
	Radius float64 // Radius of reflex circle
}

type Match struct {
	ID string // Unique ID
	Valve string // Left and right valve labels
	FromAnchor, ToAnchor Vector // Left and right anchor points
	FromTangent, ToTangent Vector // Left and right tangents
}

func Compute(uc *understand.Circuit) *Circuit {
	c := &Circuit{Name: uc.Name()}

	// Peers
	var z float64 // Total weight
	var i int
	inv := make(map[*understand.Peer]int)
	for _, p := range uc.Peer {
		inv[p] = i
		// weight := float64(len(p.Valve))
		z += 1
		c.Peer = append(c.Peer,
			&Peer{
				ID: fmt.Sprintf("peer-%s", p.Name),
				Name: template.HTMLEscapeString(fmt.Sprintf("%v", p.Name)),
				Design: template.HTMLEscapeString(fmt.Sprintf("%v", p.Design)),
				Weight: 1, //weight,
			},
		)
		i++
	}
	var u float64
	const MaxRadius = 0.9
	for _, p := range c.Peer {
		p.Angle = 2 * math.Pi * (u + p.Weight / 2) / z
		p.Anchor = CirclePointOfAngle(p.Angle)
		p.Radius = MaxRadius * p.Weight / z
		u += p.Weight
	}

	// Matchings
	for _, p := range uc.Peer { // From
		pp := c.Peer[inv[p]]
		for _, v := range p.Valve { // To
			qq := c.Peer[inv[v.Matching.Of]]
			x := pp.Anchor
			y := qq.Anchor
			c.Match = append(c.Match,
				&Match{
					ID: fmt.Sprintf("match-%s-%s", pp.Name, v.Name),
					FromAnchor: x,
					ToAnchor: y,
					FromTangent: Scalar(0.5, x),
					ToTangent:  Scalar(0.5, y),
					Valve: v.Name,
				},
			)
		}
	}
	return c
}

func CirclePointOfAngle(angle float64) Vector {
	sin, cos := math.Sincos(angle)
	return Vector{X: sin, Y: cos}
}

func Scalar(a float64, v Vector) Vector {
	return Vector{X: a*v.X, Y: a*v.Y}
}

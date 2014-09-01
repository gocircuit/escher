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

	// "github.com/gocircuit/escher/faculty"
	// . "github.com/gocircuit/escher/image"
	// "github.com/gocircuit/escher/be"
	// "github.com/gocircuit/escher/kit/plumb"
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
	Name, Design string
	Degree float64
	Angle float64 // Angle of origin-center line in [0,360]
	Radius float64 // Radius of reflex circle
}

type Match struct {
	Valve string // Left and right valve labels
	FromAnchor, ToAnchor Vector // Left and right anchor points
	FromTangent, ToTangent Vector // Left and right tangents
}

func Compute(uc *understand.Circuit) *Circuit {
	c := &Circuit{Name: uc.Name}

	// Peers
	var w float64
	var i int
	inv := make(map[*understand.Peer]int)
	for _, p := range uc.Peer {
		inv[p] = i
		deg := float64(len(p.Valve))
		w += deg
		c.Peer = append(c.Peer,
			&Peer{
				Name: fmt.Sprintf("%v", p.Name),
				Design: fmt.Sprintf("%v", p.Design),
				Degree: deg,
			},
		)
		i++
	}
	var u float64
	const MaxRadius = 0.4
	for i, p := range c.Peer {
		c.Peer[i].Angle = 2 * math.Pi * u / w
		c.Peer[i].Radius = MaxRadius * u / w
		u += p.Degree
	}

	// Matchings
	for _, p := range uc.Peer {
		pp := c.Peer[inv[p]]
		for _, v := range p.Valve {
			qq := c.Peer[inv[v.Matching.Of]]
			x := CirclePointOfAngle(pp.Angle)
			y := CirclePointOfAngle(qq.Angle)
			c.Match = append(c.Match,
				&Match{
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

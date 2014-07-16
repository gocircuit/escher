# Escher in language

(Or “HTML and JavaScript for Web 3.0 all-in-one”.)

[![Build Status](https://travis-ci.org/gocircuit/escher.svg?branch=master)](https://travis-ci.org/gocircuit/escher/escher) [![GoDoc](https://godoc.org/github.com/gocircuit/escher?status.png)](https://godoc.org/github.com/gocircuit/escher)

![Escher in language](https://github.com/gocircuit/escher/raw/master/misc/img/main.png)

Escher is a progrmaming language for everything. It can naturally represent both process and data,
while being simpler than a calculator grammar.

Escher is a language for building intelligent real-time translations between the semantics of
different physical devices accessible through chains or networks of digital or electrical technologies.

Some of the application domains of Escher are:

* Definition and generation of synthetic worlds governed by Physical laws, as in Augmented Reality and the Gaming Industry,
* General purpose concurrent and distributed programming, such as Internet services and cloud applications,
* Relational data representation, as in databases and CAD file formats,
* Real-time control loops, as in Robotics,
* Numerical and scientific computation pipelines,
* Anything that filters information from some input sources, in real-time, and sends a transformed form to output devices!

An early “proposal” for the design of Escher, 
[Escher: A black-and-white language for data and process representation](http://www.maymounkov.org/memex/abstract),
might be an informative (but not necessary) read for the theoretically inclined.

## Attention: Mathematics

The Escher abstraction of the world is NOT Turing-compatible: From the point-of-view of an
Escher program, there is no input and output: There are only emergences and disappearances of events.

Escher presents the world in a model called [Choiceless Computation](http://arxiv.org/pdf/math/9705225.pdf),
introduced by the legendary Mathematicians
[Andreas Blass](http://www.math.lsa.umich.edu/~ablass/), 
[Yuri Gurevich](http://research.microsoft.com/en-us/um/people/gurevich/) and 
[Sharon Shelah](http://shelah.logic.at/), and introduced to me by the dare-to-be-great, soon-to-be-legendary,
although-already-should-be [Benjamin Rossman](http://research.nii.ac.jp/~rossman/).

Understanding the difference between Turing Machines and Choiceless Computation,
while not entirely necessary, sheds much light on the profound difference between 
Escher and other languages. The relevant publications are quoted in the 
[bibliography](https://github.com/gocircuit/escher#bibliography) at the end.

### Choiceless programming by example

The following puzzle demonstrates choiceless programming via a simple relatable high-school 
Math puzzle:

> Four beer caps are placed on the corners of a square table with arbitrary
> orientations. There is a robot on the table that acts upon three commands:
> (a) “flip a corner cap”,
> (b) “flip two diagonal caps” and
> (c) “flip two caps along a side”

> Upon action there is no guarantee as to which corner, diagonal
> or side, respectively, the robot will choose to flip.

> Devise a sequence of commands that forces the robot to turn all caps in a
> conﬁguration where they all have the same orientation.

> Can you devise a sequence that ensures they all face up? Down?

I was introduced to the notion of choiceless programming via this puzzle by Benjamin Rossman.

## Quick start ##

Escher is an interpreter, comprising a singular executable binary. It can be built for Linux, Darwin and Windows.

Given that the [Go Language](http://golang.org) compiler is [installed](http://golang.org/doc/install),
you can build and install the circuit binary with one line:

	go install github.com/gocircuit/escher/escher

Go to the Escher base directory and run one of the tutorials

	escher -src tutorial/helloworld

## Meaning

An Escher program is a collection of interconnected _reflexes_. A reflex, the only
abstraction in Escher, represents an independent computing entity that can interact
with the “outside world” through a collection of named _valves_.

The illustration below shows a reflex, named `AND`, which has three valves,
named `X`, `Y` and `XandY`, respectively.

![An Escher reflex](https://github.com/gocircuit/escher/raw/master/misc/img/design.png)

A reflex can be implemented in another technology (currently only the 
[Go Programming Language](http://golang.org) is supported
as an external technology) or it can be composed of pre-existing reflexes.
The former is called a _gate_, while the latter is called a _circuit_.

## Gates

Gates are reflexes whose behvaior is implemented in a the underlying technology,
which is the [Go language](http://golang.org/doc/effective_go.html), at the moment.
From Escher's point-of-view (POV), gates are simply
reflexes that broker values. But from the user's POV, gates can have “side-effects”
in the “outside world” and, vice-versa, the outside world can prompt reflexive
action, such as sending out a message over a valve asynchronously.

To implement your own gates, take example from the [implementation of the 
“reasoning” gate](https://github.com/gocircuit/escher/blob/master/faculty/basic/reason.go) (discussed later).

## Circuits

Circuits are a composition of a few reflexes. 

![Boolean “not and”](https://github.com/gocircuit/escher/raw/master/misc/img/circuit.png)

Programmatically, this circuit is defined by the code:

	nand {
		// reflex recollections
		and and
		not not
		// connections
		not.X = and.XandY
		XnandY = not.notX
		and.X = X
		and.Y = Y
	}

## Syntax (files) and faculties (directories) structure

Escher programs are designated by a local root directory and all its descendants.
That directory is represented as the root in the faculty namespace inside the Escher programming environment.

Escher compiles all files ending in `.escher` and attaches the resulting circuit designs
to the namespaces corresponding to their directory parents.

To materialize (i.e. run) an Escher program, use the mandatory `-src` flag to specify the path to the local 
source directory.

	escher -src tutorial/helloworld

Escher materializes the circuit design named `main` in the root source directory, e.g.

	// The main circuit is always the one materialized (executed).
	main {
		s @show
		s = "¡Hello, world!"
	}

## Basic gates

By default, the Escher environment provides a basic set of gates (a basis),
which enable a rich (infinite) language of possibilities in data manipulation.

Collectively, they are data (concept) and transformation (sentence) gates

These gates are not part of Escher's semantics. They are merely an optional
library—a playground for beginners. Users can implement their own gates
for data and transformation.

The basis reference below is nearly entirely visual. You will notice that the
visual language follows a prescribed format.

### Data (Noun) gates

On blank slate, there is “nothing to do”—so to speak. For this reason,
we have a collection of gates which are effectively “springs” of objects.

Some produce integers, some floats, some complex numbers, some strings.
These are familiar types.

There is one gate that produces “trees”. Trees are the basic type of
“weavable” (or mutable) knowledge. (In fact, the other types are not
necessary but we throw them in for convenience.)

In the illustration below, the syntax of the respective gate design is displayed
as their name (white on black).

![Impression of the mind](https://github.com/gocircuit/escher/raw/master/misc/img/impress.png)

### Combinator (Manipulator) gates

![Grammar manipulation gates](https://github.com/gocircuit/escher/raw/master/misc/img/combine.png)

### Arithmetic (Applying) gates

Arithmetic gates are a sufficient basis of operations that enables
algorithmic manipulation of the types string, int, float and complex.
Coming soon.

### Tree of Knowledge and the Reason (Learning) Gate

You will notice, one of the basic data gates allows the creation of a struct-like object.
This is called a _tree_. It is a novel data structure, described in detail in
[“Extensible records with scoped labels” by Daan Leijen](http://research.microsoft.com/pubs/65409/scopedlabels.pdf).

These data structures are “built out” and “trimmed down” using three elegant
reflection methods, described in the above publication. 

Escher embodies all three in one gate, whose main purpose is to manipulate
the contents of trees. This is the _Reason Gate_, illustrated below.

The following three illustrations show the same gate design, but under different
orientations of the event streams. In all valid cases, the relationship
between the valve values shown examplifies the effect of the gate.

![Generalization](https://github.com/gocircuit/escher/raw/master/misc/img/generalization.png)

> Belief for the current state of the world, combined with a new observation, results in
a theory.

![Explanation](https://github.com/gocircuit/escher/raw/master/misc/img/explanation.png)

> A theory of observations, which explains (includes) an observation at hand, explains
the observation only to a belief consisting of the theory without the observation.

![Prediction](https://github.com/gocircuit/escher/raw/master/misc/img/prediction.png)

> When a belief of the state of the world is combined with a theory that is bigger,
the conjectured difference must be found in a new observation.

## Duality gates

Duality gates are the boundary between Escher semantics and the outside world. 
They are the I/O with the outside. Such gates affect some external technology
when prompted through Escher in a certain way. Alternatively, such gates might
fire an Escher message on one of its valves, in response to an asynchronous
events occuring in an external technology.

### Variation (Surprise) and Causation (Action) gates

![See and Show](https://github.com/gocircuit/escher/raw/master/misc/img/seeshow.png)

For instance, with the gates we've seen so far, one might construct the following higher-level
circuit abstraction for an I/O device, which is controlled by a defered logic:

![I/O device](https://github.com/gocircuit/escher/raw/master/misc/img/io.png)

And the respective source code:

	io_device {
		// recalls
		in see
		out show
		swtch switch
		// matchings
		Logic = swtch.Socialize
		in.Sensation = swtch.Hear
		out.Action = swtch.Speak
	}

## Introspective and extrospective gates

This special type of gates fulfills the complementary functions of
constructing new circuit designs “dynamically” (akin to “reflection” in other languages),
and materializing (i.e. executing) these designs. Coming soon.

### The Julia (Exploiting) Gate

Coming soon.

### The Escher (Teaching) Gate

Coming soon.

## The future collapsed

I envision that in the natural course of action at play, … (more coming soon).

## And…

…if you think this language is `#KingOfMetaphor`, please, tweet to
[@StephenAtHome](https://twitter.com/StephenAtHome) that his title of
`#KingOfMetaphor` is being challenged, in a good way. Tweet
that `@escherio` wants a `#ColbertBump`. The `#ColbertFaculty` is coming soon.

…if you want to inquire about the science behind [@escherio](https://twitter.com/escherio), tweet to me,
Petar [@maymounkov](https://twitter.com/maymounkov).

Or, simply lose yourself in the 
[initial](http://www.maymounkov.org/chomsky-valiant-algorithmic-mirror)
[thoughts](http://www.maymounkov.org/puzzle-test-turing-test) that
led to the invention of Escher.

To me, Escher is a language for weaving dreams: It makes imagination real. Help me make it tangible, so it can be shared.

## Sponsors and credits

* [DARPA XDATA](http://www.darpa.mil/Our_Work/I2O/Programs/XDATA.aspx) initiative 
under the program management of
[Dr. Christopher White](http://www.darpa.mil/Our_Work/I2O/Personnel/Dr_Christopher_White.aspx), 2013–2014
* [Data Tactics Corp](http://www.data-tactics.com/), 2013–2014
* [L3](http://www.l-3com.com/), 2014

## Bibliography

* [Choiceless Polynomial Time](http://arxiv.org/pdf/math/9705225.pdf), 
[Andreas Blass](http://www.math.lsa.umich.edu/~ablass/), 
[Yuri Gurevich](http://research.microsoft.com/en-us/um/people/gurevich/) and 
[Saharon Shelah](http://shelah.logic.at/), published by Shelah Office
on [arXiv](http://arxiv.org/abs/math/9705225), 1997

* [Choiceless Polynomial Time Logic: Inability to Express](http://link.springer.com/chapter/10.1007%2F3-540-44622-2_6),
[Saharon Shelah](http://shelah.logic.at/), Springer Lecture Notes in Computer Science, Volume 1862, 2000, pp. 72-125

* [Circuits of the Mind](http://www.amazon.com/Circuits-Mind-Leslie-G-Valiant/dp/0195126688), Leslie Valiant, 2000
* [Probably Approximately Correct, 53589083](http://www.probablyapproximatelycorrect.com/), Leslie Valiant, 2013

Additional, enlightening reading includes:

* [Ancient Evenings](http://en.wikipedia.org/wiki/Ancient_Evenings), Norman Mailer, 1983
* [Yoga Philosophy of Patañjali](http://www.amazon.com/Yoga-Philosophy-Patanjali-Translation-Annotations/dp/0873957296),
State University of New York Press, 1984

# Escher

[![Build Status](https://travis-ci.org/gocircuit/escher.svg?branch=master)](https://travis-ci.org/gocircuit/escher/escher) [![GoDoc](https://godoc.org/github.com/gocircuit/escher?status.png)](https://godoc.org/github.com/gocircuit/escher)

Escher is a progrmaming language for everything. It can naturally represent both process and data,
while being simpler than a calculator. The remainder of this page constitutes a complete documentation.

Escher is a language for building intelligent real-time translations between the semantics of
different physical devices accessible through chains or networks of digital or electrical technologies.

Some of the application domains of Escher are:

* Definition and generation of synthetic worlds governed by Physical laws, as in Augmented Reality and the Gaming Industry,
* General purpose concurrent and distributed programming, such as Internet services and cloud applications,
* Relational data representation, as in databases and CAD file formats,
* Real-time control loops, as in Robotics,
* Numerical and scientific computation pipelines,
* And so on.

An early “proposal” for the design of Escher, 
[Escher: A black-and-white language for data and process representation](http://www.maymounkov.org/memex/abstract),
might be an informative (but not necessary) read for the theoretically inclined.

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

??

## Circuits

Circuits are a composition of a few reflexes. 

![Boolean “not and”](https://github.com/gocircuit/escher/raw/master/misc/img/circuit.png)

Programmatically, this gate is expressed as

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

	// The main circuit is always the one materialized (executed).
	main {
		s @show
		s = "¡Hello, world!"
	}

## Data  (Concept) and transformation (Sentence) gates

### Data (Noun) gates

![Impression of the mind](https://github.com/gocircuit/escher/raw/master/misc/img/impress.png)

### Combinator (Manipulator) gates

![Grammar manipulation gates](https://github.com/gocircuit/escher/raw/master/misc/img/combine.png)

### Arithmetic (Applying) gates

Coming soon.

### The Reason (Learning) Gate

![Generalization](https://github.com/gocircuit/escher/raw/master/misc/img/generalization.png)

![Explanation](https://github.com/gocircuit/escher/raw/master/misc/img/explanation.png)

![Prediction](https://github.com/gocircuit/escher/raw/master/misc/img/prediction.png)

## Duality gates

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

…or, lose yourself in the 
[initial](http://www.maymounkov.org/chomsky-valiant-algorithmic-mirror)
[thoughts](http://www.maymounkov.org/puzzle-test-turing-test) that
led to the invention of Escher.

## Sponsors and credits

* [DARPA XDATA](http://www.darpa.mil/Our_Work/I2O/Programs/XDATA.aspx) initiative, 2013–2014
* [Data Tactics Corp](http://www.data-tactics.com/), 2013-2014
* [L3](http://www.l-3com.com/), 2014

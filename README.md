# Escher

Escher is a progrmaming language for everything. It can naturally represent both process and data,
while being simpler than a calculator. The remainder of this page constitutes a complete documentation.

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

## Circuits

Circuits are a composition of a few reflexes. 

![Boolean “not and”](https://github.com/gocircuit/escher/raw/master/misc/img/circuit.png)


## Data and transformation gates

### Arithmetic (Application) gates

### The Reasoning (Learning) Gate

![Generalization](https://github.com/gocircuit/escher/raw/master/misc/img/generalization.png)

![Explanation](https://github.com/gocircuit/escher/raw/master/misc/img/explanation.png)

![Prediction](https://github.com/gocircuit/escher/raw/master/misc/img/prediction.png)

## Introspective and extrospective gates

## The Fractal (Exploiting) Gate

Coming soon.

## The Escher (Teaching) Gate

Coming soon.

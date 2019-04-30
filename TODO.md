# TODO - Escher language

* tool to discover blockages
* rename verb to directive
* escher.Replace gate to substitute the residual of containing circuit with…
* addresses are sugar for a two-sided reflex:
  syntactic address and index on one side, and ...
	* generalize
* download wikipedia dataset
* file reader materializer
* convert non-escher files in source directory in materializers of respective file readers

## THINK

* remove name/value distinction (delayed because go map keys cannot be circuit or other non-primitves at the moment)
	* possible resolution: make all Go circuit manipulations functional
* (maybe) convert all these todos in here to issues (on github)
* device some standard for storing a set of attributes with each gate,
  which can be used for graphical representation:
```escher
myCircuit {
	gateA 123            `// @visual{ colorFg { 255; 0; 0; }; colorBg { 0; 0; 0; }; position2d { 1.0; 1.0; }; position3d { 1.0; 1.0; 1.0; }; }`
	gateB `some value`   `// Some textual comment here @visual{ colorFg { 255; 255; 255; 255; }; colorBg { 0; 0; 0; 255; }; position2d { 1.0; 2.0; }; position3d { 1.0; 2.0; 3.0; }; }`
}
```
* device and formulate a standard naming convention,
  to be used at least in the Esche repo
  (think Java: `ClassName`, `getSomething()`, `CONSTANT_VALUE`, ...)


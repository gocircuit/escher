// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package browser

type Msg struct {
	ID  string `json:"id"`
	Pay string `json:"pay"`
}

type Inject struct {
	Func string      `json:"func"` // Func is "function(id, arg) { … }"
	Arg  interface{} `json:"arg"`  // Arg is a JSON encoding of an Arg instance
}

// type Arg map[string]interface{} // name –> arg

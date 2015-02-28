package weaver

import (
	"reflect"
)

// type UserReflex struct {
//	*Reflex
// 	Q    struct {
// 		A int
// 		B Weaver
// 	}
// 	R struct {
// 		R float
// 	}
// }

// func (r *UserReflex) Spark() {
// 	// …
// }

// Rule
type Rule struct {
	receiver reflect.Value
}

func NewRule(rule interface{}) Rule {
	//??
}

func (r *Rule) Sources() []Name {
	//xx
}

func (r *Rule) Sinks() []Name {
	//xx
}

func (r *Rule) Write(Name, Value) {
	//xx
}

func (r *Rule) Spark() {
	//xx
}

func (r *Rule) Read(Name) Value {
	//xx
}

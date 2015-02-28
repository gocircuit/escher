package weaver

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

type Rule interface {
	Sources() []Name
	Sinks() []Name
	Write(Name, Value)
	Spark()
	Read(Name) Value
}

// // Rule
// type rule struct {
// 	receiver reflect.Value
// }

// func NewRule(rule interface{}) Rule {
// 	//??
// }

// func (r *rule) Sources() []Name {
// 	//xx
// }

// func (r *rule) Sinks() []Name {
// 	//xx
// }

// func (r *rule) Write(Name, Value) {
// 	//xx
// }

// func (r *rule) Spark() {
// 	//xx
// }

// func (r *rule) Read(Name) Value {
// 	//xx
// }

package graphlator

type Parameter struct {
	Constraint constraint
	Predicate  string
	Value      interface{}
}

type Result struct {
	Name string
	// Filter Filter
	Result *Result
}

type Function struct {
	Name      string
	Parameter *Parameter
	Result    *Result
}

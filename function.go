package graphlator

type Parameter struct {
	Operation operation
	Predicate string
	Value     interface{}
}

func (p Parameter) AsFilter() *Filters {
	return &Filters{
		constraint: -1,
		params:     []Parameter{p},
	}
}

type Function struct {
	Name      string
	Parameter *Parameter
	Filters   *Filters
	Results   []Result
}

type Result struct {
	want       string
	isExpanded bool
	Expanded   []Result
}

func ResultSlice(r ...Result) []Result {
	return r
}

func NewResult(want string, expanded ...Result) Result {
	r := Result{want: want}

	if len(expanded) != 0 {
		r.isExpanded = true
		r.Expanded = expanded
	}

	return r
}

type Filters struct {
	constraint constraint
	params     []Parameter
}

func ParamToFilter(p Parameter) *Filters {
	return &Filters{
		constraint: -1,
		params:     []Parameter{p},
	}
}

func And(p ...Parameter) *Filters {
	return &Filters{
		constraint: and,
		params:     p,
	}
}

func Or(p ...Parameter) *Filters {
	return &Filters{
		constraint: or,
		params:     p,
	}
}

func Not(p ...Parameter) *Filters {
	return &Filters{
		constraint: not,
		params:     p,
	}
}

package graphlator

type Parameter struct {
	Operation operation
	Predicate string
	Value     interface{}
}

func (p *Parameter) AsFilter() *Filters {
	return &Filters{
		constraint: -1,
		params:     []*Parameter{p},
	}
}

type Result struct {
	Name string
	// Filter Filter
	Result *Result
}

type Function struct {
	Name      string
	Parameter *Parameter
	Filters   *Filters
	Result    *Result
}

type Filters struct {
	constraint constraint
	params     []*Parameter
}

func ParamToFilter(p *Parameter) *Filters {
	return &Filters{
		constraint: -1,
		params:     []*Parameter{p},
	}
}

func And(p ...*Parameter) *Filters {
	return &Filters{
		constraint: and,
		params:     p,
	}
}

func Or(p ...*Parameter) *Filters {
	return &Filters{
		constraint: or,
		params:     p,
	}
}

func Not(p ...*Parameter) *Filters {
	return &Filters{
		constraint: not,
		params:     p,
	}
}

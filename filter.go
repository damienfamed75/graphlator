package graphlator

type Filters struct {
	constraint constraint
	params     []Parameter
}

func (f *Filters) UpdateConstraint(c constraint) {
	f.constraint = c
}

func ParamToFilter(p Parameter) *Filters {
	return &Filters{
		constraint: invalidConstraint,
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

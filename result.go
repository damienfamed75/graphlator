package graphlator

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

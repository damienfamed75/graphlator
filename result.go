package graphlator

type Result struct {
	want       string
	isExpanded bool
	ResultSlice
}

func NewResult(want string, expanded ...Result) Result {
	r := Result{want: want}

	if len(expanded) != 0 {
		r.isExpanded = true
		r.ResultSlice = expanded
	}

	return r
}

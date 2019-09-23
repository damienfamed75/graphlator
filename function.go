package graphlator

type Function struct {
	Name      string
	Parameter Parameter
	Filters   *Filters
	Results   []Result
}

func (f *Function) InsertResults(results ...Result) {
	f.Results = append(f.Results, results...)
}

func (f *Function) UpsertResults(results ...Result) {
	for _, res := range results {
		if !f.ResultExists(res) {
			f.Results = append(f.Results, res)
		}
	}
}

func (f *Function) ResultExists(result Result) bool {
	for _, r := range f.Results {
		if r.want == result.want {
			return true
		}
	}

	return false
}

func (f *Function) RemoveResult(want string) []Result {
	var tmp []Result

	for _, r := range f.Results {
		if r.want == want {
			continue
		}

		tmp = append(tmp, r)
	}

	f.Results = tmp

	return tmp
}

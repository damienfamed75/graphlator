package graphlator

// ResultSlice is the type used for functions and Results themselves.
// This type exists to simplify the code base for manipulating slices of results
// instead of having duplicate code for the Function type and Result type.
type ResultSlice []Result

// Insert will insert any Result regardless if it already exists in the slice
// or not.
func (r *ResultSlice) Insert(res ...Result) {
	*r = append(*r, res...)
}

// Upsert will add only new results given. Any pre-existing results
// will be omitted and then the integer result is the number of results
// that ended up being added to the slice.
func (r *ResultSlice) Upsert(res ...Result) int {
	var added int

	for _, rr := range res {
		if !r.Exists(rr.want) {
			added++
			*r = append(*r, rr)
		}
	}

	return added
}

// Exists is just a simple conditional that will return a boolean value
// instead of a reference of a Result unlike Find.
func (r ResultSlice) Exists(want string) bool {
	for _, res := range r {
		if res.want == want {
			return true
		}
	}

	return false
}

// IsEmpty is a conditional to see if the result slice has no children.
func (r ResultSlice) IsEmpty() bool {
	return len(r) == 0
}

// Find returns a reference to a result that contains the matching string.
func (r ResultSlice) Find(want string) *Result {
	for _, res := range r {
		if res.want == want {
			return &res
		}
	}

	return nil
}

// Remove will find any surface level results and slice out the found
// result with the matching string.
func (r *ResultSlice) Remove(want string) bool {
	if r.IsEmpty() {
		return false
	}

	var found bool
	var index int

	for i, res := range *r {
		if res.want == want {
			index = i
			found = true
			break
		}
	}

	if found {
		results := *r

		if len(*r) <= index+1 {
			results = results[:index]
		} else {
			results = append(results[:index], results[index+1:]...)
		}

		*r = results
		return true
	}

	return false
}

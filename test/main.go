package main

import (
	"fmt"

	. "github.com/damienfamed75/graphlator"
)

// Types of Results
//
// name uid
//
// has { name }
//
// has @filter(type(Friend)) { name }
//
// expand(_all_)
//
// expand has {
//
// }

func main() {
	f := Function{
		Name:      "me",
		Parameter: Equal("name", "Damien"),
		Filters:   GreaterThan("age", 18).AsFilter(),
		Result: &Result{
			Name: "age",
		},
		// Get multiple results back.
	}

	t := NewTranslator(WithLanguage(GraphQLPlus))

	res := t.TranslateQuery(f)

	fmt.Printf("%s\n", res)
}

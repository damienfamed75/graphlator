package main

import (
	"fmt"

	. "github.com/damienfamed75/graphlator"
)

func main() {
	f := Function{
		Name:      "me",
		Parameter: Equal("name", "Damien"),
		Filters:   GreaterThan("age", 18).AsFilter(),
		Results: ResultSlice(
			NewResult("name"), NewResult("uid"),
			NewResult("has", NewResult("name"), NewResult("uid")),
		),
	}

	t := NewTranslator(WithLanguage(GraphQLPlus))

	res := t.TranslateQuery(f)

	fmt.Printf("%s\n", res)
}

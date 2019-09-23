package main

import (
	"fmt"

	. "github.com/damienfamed75/graphlator"
)

func main() {
	f := Function{
		Name:      "me",
		Parameter: UID("0x123"),
		Filters:   And(GreaterThan("age", 18), Type("Friend")),
		Results: ResultSlice(
			NewResult("name"), NewResult("uid"),
			NewResult("has", NewResult("name"), NewResult("uid")),
		),
	}

	t := NewTranslator(WithLanguage(GraphQLPlus))

	res := t.TranslateQuery(f)

	fmt.Printf("%s\n", res)
}

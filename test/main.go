package main

import (
	"fmt"

	g "github.com/damienfamed75/graphlator"
)

func main() {
	f := g.Function{
		Name:      "me",
		Parameter: g.Equal("name", "Damien"),
		Result: &g.Result{
			Name: "name",
		},
	}

	t := g.NewTranslator(g.WithLanguage(g.GraphQLPlus))

	res := t.TranslateQuery(f)

	fmt.Printf("%s\n", res)
}

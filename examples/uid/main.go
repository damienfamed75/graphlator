package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	gl "github.com/damienfamed75/graphlator"
	"github.com/damienfamed75/quirk/v2"
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

// Person is the structure to hold the node's data.
// When using quirk you must have tags associated with your fields.
// The first quirk parameter is the name of the predicate in Dgraph.
// The second parameter (which always is "unique") specifies if this
// field should be unique throughout the graph.
type Person struct {
	Type   string `quirk:"dgraph.type"`
	Name   string `quirk:"name"`
	SSN    string `quirk:"ssn,unique"`
	Policy string `quirk:"policy,unique"`
}

func main() {
	conn, err := grpc.Dial("127.0.0.1:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed when dialing grpc: %v", err)
	}
	defer conn.Close()

	dg := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	// Set schema and add nodes to Dgraph to query on.
	uid, err := setupTestNodes(dg)
	if err != nil {
		log.Fatalf("error setting up dgraph: %v\n", err)
	}

	// Create a new translator for the queries.
	t := gl.NewTranslator(gl.WithLanguage(gl.GraphQLPlus))
	// Translate a structured format to a string query for dgraph.
	res := t.TranslateQuery(
		gl.Function{
			Name:      "me",
			Parameter: gl.UID(uid),
			Results: gl.ResultSlice(
				gl.NewResult("name"), gl.NewResult("uid"),
			),
		},
	)

	// Use the dgo client to query for the nodes with our resulted query string.
	txn := dg.NewTxn()
	assigned, err := txn.Query(context.Background(), res)
	if err != nil {
		log.Fatalf("error querying for nodes: %v", err)
	}

	fmt.Printf("Query Result: %s\n", assigned.GetJson())
}

func setupTestNodes(dg *dgo.Dgraph) (string, error) {
	err := dg.Alter(context.Background(), &api.Operation{
		Schema: `
		name: string @index(hash) .
		ssn: string @index(hash) @upsert .
		policy: string @index(hash) @upsert .

		type person {
			name: string
			ssn: string
			policy: string
		}
		`,
	})
	if err != nil {
		return "", fmt.Errorf("alteration error with setting schema: %w", err)
	}

	// Upsert the node of type "person".
	c := quirk.NewClient()
	uidMap, err := c.InsertNode(context.Background(), dg, &quirk.Operation{
		SetSingleStruct: &Person{Type: "person", Name: "Damien", SSN: "123", Policy: "JKL"},
	})
	if err != nil {
		return "", fmt.Errorf("error inserting node: %w", err)
	}

	// If we didn't get a resulting UID back then error out.
	if _, ok := uidMap["Damien"]; !ok {
		return "", errors.New("UID map returned empty")
	}

	return uidMap["Damien"].Value(), nil
}

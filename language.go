package graphlator

// Language determines what kind of result you will get back from the translator.
type Language int8

const (
	// GraphQL assigns the translator to turn structured formatted functions
	// into a GraphQL string query.
	// Note: This functionality is not begun.
	GraphQL Language = iota
	// GraphQLPlus assigned the translator to output Dgraph's GraphQL+- language
	// for its string queries.
	// Note: This functionality is not yet finished.
	GraphQLPlus
)

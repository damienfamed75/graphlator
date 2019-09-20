package graphlator

type Language int8

const (
	GraphQL Language = iota
	GraphQLPlus
)

// go:generate stringer -type=constraint -linecomment
type constraint int8

// They have a prefix of constraint because "type" isn't a valid const name.
const (
	constraintEqual              constraint = iota // eq
	constraintGreaterThanOrEqual                   // ge
	constraintLessThanOrEqual                      // le
	constraintGreaterThan                          // gt
	constraintLessThan                             // lt
	constraintAnyOfTerms                           // anyOfTerms
	constraintAllOfTerms                           // allOfTerms
	constraintType
)

func Equal(predicate string, value interface{}) *Parameter {
	return &Parameter{
		Constraint: constraintEqual,
		Predicate:  predicate,
		Value:      value,
	}
}

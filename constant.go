package graphlator

type Language int8

const (
	GraphQL Language = iota
	GraphQLPlus
)

//go:generate stringer -type=constraint -linecomment
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
	constraintType                                 // type
)

func Equal(predicate string, value interface{}) *Parameter {
	return buildParamater(constraintEqual, predicate, value)
}

func GreaterThanOrEqual(predicate string, value interface{}) *Parameter {
	return buildParamater(constraintGreaterThanOrEqual, predicate, value)
}

func LessThanOrEqual(predicate string, value interface{}) *Parameter {
	return buildParamater(constraintLessThanOrEqual, predicate, value)
}

func GreaterThan(predicate string, value interface{}) *Parameter {
	return buildParamater(constraintGreaterThan, predicate, value)
}

func LessThan(predicate string, value interface{}) *Parameter {
	return buildParamater(constraintLessThan, predicate, value)
}

func AnyOfTerms(predicate string, value interface{}) *Parameter {
	return buildParamater(constraintAnyOfTerms, predicate, value)
}

func AllOfTerms(predicate string, value interface{}) *Parameter {
	return buildParamater(constraintAllOfTerms, predicate, value)
}

func Type(typeAsString string) *Parameter {
	return buildParamater(constraintType, typeAsString, nil)
}

func buildParamater(c constraint, p string, v interface{}) *Parameter {
	return &Parameter{
		Constraint: c,
		Predicate:  p,
		Value:      v,
	}
}

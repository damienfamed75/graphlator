package graphlator

//go:generate stringer -type=operation -linecomment
type operation int8

// They have a prefix of operation because "type" isn't a valid const name.
const (
	operationEqual              operation = iota // eq
	operationGreaterThanOrEqual                  // ge
	operationLessThanOrEqual                     // le
	operationGreaterThan                         // gt
	operationLessThan                            // lt
	operationAnyOfTerms                          // anyOfTerms
	operationAllOfTerms                          // allOfTerms
	operationType                                // type
	operationHas                                 // has
	operationRegexp                              // regexp
	operationUIDIn                               // uid_in
	operationUID                                 // uid
)

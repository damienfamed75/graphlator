package graphlator

//go:generate stringer -type=constraint -linecomment
type constraint int8

const (
	invalidConstraint constraint = iota - 1 // INVALID
	and                                     // AND
	or                                      // OR
	not                                     // NOT
)

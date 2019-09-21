package graphlator

//go:generate stringer -type=constraint -linecomment
type constraint int8

const (
	and constraint = iota // AND
	or                    // OR
	not                   // NOT
)

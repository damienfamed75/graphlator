package graphlator

import (
	"fmt"
	"strings"
	"sync"
)

type Translator struct {
	DefaultLanguage Language
	buffers         sync.Pool
	adjustments     option
}

func setupTranslator() *Translator {
	return &Translator{
		DefaultLanguage: GraphQLPlus,
		buffers: sync.Pool{
			New: func() interface{} {
				return &strings.Builder{}
			},
		},
	}
}

func NewTranslator(r ...Regulation) *Translator {
	t := setupTranslator()

	for _, reg := range r {
		reg(&t.adjustments)
	}

	return t
}

func (t *Translator) TranslateQuery(f Function, r ...Regulation) string {
	buf := t.buffers.Get().(*strings.Builder)
	defer func() {
		buf.Reset()
		t.buffers.Put(buf)
	}()

	buf.WriteString(fmt.Sprintf("{%s(func: %s(",
		f.Name, f.Parameter.Operation))

	writeOperation(buf, f.Parameter)
	buf.WriteString(") ")

	if f.Filters != nil {
		buf.WriteString("@filter(")
		for i, p := range f.Filters.params {
			if f.Filters.constraint != invalidConstraint && i != 0 {
				buf.WriteString(" " + f.Filters.constraint.String() + " ")
			}
			buf.WriteString(p.Operation.String() + "(")
			writeOperation(buf, p)

		}
		buf.WriteByte(')')
	}

	loopResults(buf, f.Results)

	buf.WriteByte('}') // End of Query

	query := buf.String()

	return query
}

func writeOperation(buf StringByteWriter, p Parameter) {
	switch p.Operation {
	case operationUIDIn:
		buf.WriteString(fmt.Sprintf("%s,%v)", p.Predicate, p.Value))
		fallthrough
	case operationType:
		buf.WriteString("\"" + p.Predicate + "\"")
	default:
		buf.WriteString(p.Predicate)
	}

	if p.Value != nil {
		buf.WriteString(fmt.Sprintf(",%#v", p.Value))
	}
	buf.WriteByte(')')
}

func loopResults(buf StringByteWriter, r []Result) {
	buf.WriteByte('{')
	for _, rr := range r {
		buf.WriteString(rr.want)
		if rr.isExpanded {
			loopResults(buf, rr.Expanded)
		}
		buf.WriteByte(' ')
	}

	buf.WriteByte('}')
}

/*

GraphQL+

{
	me(func: type("person")) @filter(eq(name, "Damien")) {
		uid
		name
	}
}

GraphQL

{
	person(name: "Damien") {
		uid
		name
	}
}

*/

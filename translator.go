package graphlator

import (
	"bytes"
	"fmt"
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
				return &bytes.Buffer{}
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

func (t *Translator) TranslateQuery(f Function, r ...Regulation) []byte {
	buf := t.buffers.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		t.buffers.Put(buf)
	}()

	buf.WriteByte('{') // Beginning of Query

	buf.WriteString(f.Name + "(func: ")
	buf.WriteString(f.Parameter.Operation.String() + "(")

	// switch f.Parameter.Operation {
	// case operationType:
	// 	buf.WriteString("\"" + f.Parameter.Predicate + "\"") // TODO Change dynamic
	// default:
	// 	buf.WriteString(f.Parameter.Predicate) // TODO Change dynamic
	// }

	// if f.Parameter.Value != nil {
	// 	buf.WriteString(fmt.Sprintf(",%#v", f.Parameter.Value))
	// }
	// buf.WriteString("))")
	writeOperation(buf, f.Parameter)
	buf.WriteString(") ")

	if f.Filters != nil {
		buf.WriteString("@filter(")
		for i, p := range f.Filters.params {
			if f.Filters.constraint != -1 && i != 0 {
				buf.WriteString(" " + f.Filters.constraint.String() + " ")
			}
			buf.WriteString(p.Operation.String() + "(")
			// buf.WriteString(p.Predicate + ",") // TODO Change dynamic
			// buf.WriteString(fmt.Sprintf("%#v)", p.Value))
			writeOperation(buf, p)

		}
		buf.WriteByte(')')
	}

	loopResults(buf, f.Results)

	buf.WriteByte('}') // End of Query

	query := buf.Bytes()

	return query
}

func writeOperation(buf *bytes.Buffer, p Parameter) {
	switch p.Operation {
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

func loopResults(buf *bytes.Buffer, r []Result) {
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
	me(func: eq(name, "Damien")) {
		uid
		name
	}
}

*/

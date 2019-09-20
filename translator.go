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

	buf.WriteByte('{')

	buf.WriteString(f.Name + "(func: ")
	buf.WriteString(f.Parameter.Constraint.String() + "(")
	buf.WriteString(f.Parameter.Predicate + ",")
	buf.WriteString(fmt.Sprintf("%#v))", f.Parameter.Value))

	buf.WriteByte('{')
	buf.WriteString(f.Result.Name)
	buf.WriteByte('}')

	buf.WriteByte('}')

	query := buf.Bytes()

	buf.Reset()
	t.buffers.Put(buf)

	return query
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

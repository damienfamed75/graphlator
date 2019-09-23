package graphlator

type option struct {
	language Language
}

// Regulation is the type name for any augmentations that are made to the
// translator.
type Regulation func(o *option)

// WithLanguage changes the mode that the translator is in.
// This can be either GraphQL or GraphQLPlus
func WithLanguage(l Language) Regulation {
	return func(o *option) {
		o.language = l
	}
}

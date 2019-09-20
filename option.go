package graphlator

type option struct {
	language Language
}

type Regulation func(o *option)

func WithLanguage(l Language) Regulation {
	return func(o *option) {
		o.language = l
	}
}

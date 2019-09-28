package graphlator

type Function struct {
	Name      string
	Parameter Parameter
	Filters   *Filters
	Results   ResultSlice
}

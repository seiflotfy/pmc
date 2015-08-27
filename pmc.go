package pmc

/*
Sketch ...
*/
type Sketch struct {
	l uint
}

/*
NewSketch ...
*/
func NewSketch(l uint) *Sketch {
	return &Sketch{l: l}
}

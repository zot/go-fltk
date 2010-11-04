package fltk

type LayoutCalc struct {
	X, Y, W, H int
}

func NewLayoutCalc(w Widgety) *LayoutCalc {
	return  &LayoutCalc{0, 0, w.W(), w.H()}
}
// add a widget of h height
func (l *LayoutCalc) Add(w Widgety) {
	l.Y += w.H()
	l.H -= w.H()
}

package fltk

type LayoutCalc struct {
	X, Y, W, H int
}

func NewLayoutCalc(w Widgety) *LayoutCalc {
	return  &LayoutCalc{0, 0, w.getWidget().W(), w.getWidget().H()}
}
// add a widget of h height
func (l *LayoutCalc) Add(w Widgety) {
	l.Y += w.getWidget().H()
	l.H -= w.getWidget().H()
}

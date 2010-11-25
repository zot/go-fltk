package main

import "fltk"

func main() {
	w := fltk.NewWindow(300, 180)
	w.Begin()
	b := fltk.NewWidget(20, 40, 260, 100, "Hello there!")
	b.Box(fltk.UP_BOX)
	b.LabelFont(fltk.HELVETICA_BOLD_ITALIC)
	b.LabelSize(36)
	b.LabelType(fltk.SHADOW_LABEL)
	w.End()
	w.Show([]string{})
	fltk.Run(func() bool {return true})
}

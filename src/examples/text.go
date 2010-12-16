package main

import "fltk"
import "fmt"
import "strconv"

func main() {
	exit := false
	window := fltk.NewWindow(300, 180)
	l := fltk.NewLayoutCalc(window)
	window.SetCallback(func(){exit = true; println("CLOSED"); window.Destroy()})
	window.Begin()
	i := fltk.NewInput(l.X, l.Y, l.W, 30)
	l.Add(i)
	i.StealEvents(fltk.PUSH_MASK | fltk.DRAG_MASK | fltk.PASTE_MASK)
	i.SetEventHandler(func(e *fltk.Event) {
		if (e.Stolen) {
			fmt.Println("STOLE EVENT:", strconv.Itoa(e.Event), "widget:", e.Widget, "push: ", fltk.PUSH)
			if (e.Event == fltk.PUSH) {
				fmt.Println("\nMouse Position:", i.MousePosition(), "\n")
			}
			if (e.Event != fltk.PUSH && e.Event != fltk.RELEASE && e.Event != fltk.PASTE) || e.Button == 1 {
				fmt.Println("CONTINUING EVENT:", strconv.Itoa(e.Event), "widget:", e.Widget, "push: ", fltk.PUSH)
				e.Continue()
			}
		}
	})
	fmt.Println("Editor at y: ", l.Y)
	e := fltk.NewTextEditor(l.X, l.Y, l.W, l.H)
	window.Resizable(e)
	window.End()
	window.Show([]string{})
	fltk.Run(func() bool {return !exit})
}

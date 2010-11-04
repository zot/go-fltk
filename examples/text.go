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
	i.StealEvents(fltk.PUSH_MASK | fltk.DRAG_MASK)
	i.SetEventHandler(func(e *fltk.Event) {
		if (e.Stolen) {
			fmt.Println("CONTINUING EVENT:", strconv.Itoa(e.Event), "widget:", e.Widget)
			i.ContinueEvent()
		}
	})
	fmt.Println("Editor at y: ", l.Y)
	e := fltk.NewTextEditor(l.X, l.Y, l.W, l.H)
	window.Resizable(e)
	window.End()
	window.Show([]string{})
	/*
	for !exit {
		fltk.Wait()
	}
	 */
	println("RUN")
	fltk.Run()
}

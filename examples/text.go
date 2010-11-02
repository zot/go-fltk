package main

import "fltk"
//import "fmt"
import "strconv"

func main() {
	exit := false

	window := fltk.NewWindow(300, 180)
	window.SetCallback(func(){exit = true; println("CLOSED")})
	window.Begin()
	i := fltk.NewInput(0, 0, 300, 30)
	i.StealEvents(fltk.PUSH_MASK | fltk.DRAG_MASK)
	i.SetEventHandler(func(e *fltk.Event) {
		if (e.Stolen) {
			println("CONTINUING EVENT: " + strconv.Itoa(e.Event))
			i.ContinueEvent()
		}
	})
	window.End()
	window.Show([]string{})
	for !exit {
		fltk.Wait()
	}
}

package main

import "fltk"
import "fmt"
import "strconv"

func addText(fr *fltk.OberFrame) {
	fr.Add(fltk.NewTextEditor(fr.X(), fr.Y(), fr.W(), fr.H()))
}

func makeFrame(horizontal bool, parent *fltk.Group, x, y, w, h int) *fltk.OberFrame {
	f := fltk.NewOberFrame(x, y, w, h, horizontal)
	parent.Add(f)
	i := fltk.NewInput(f.X(), f.Y(), f.W(), 30)
	f.Add(i)
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
	return f
}

func main() {
	exit := false
	window := fltk.NewWindow(300, 180)
	window.SetCallback(func(){exit = true; println("CLOSED"); window.Destroy()})
	f := makeFrame(true, &window.Group, 0, 0, 300, 180)
	window.Resizable(f)
	g := makeFrame(false, &f.Group, 0, 0, 100, 180)
	h := makeFrame(false, &f.Group, 100, 0, 100, 180)
	i := makeFrame(false, &f.Group, 200, 0, 100, 180)
	addText(g)
	addText(h)
	addText(i)
	window.Show([]string{})
	fltk.Run(func() bool {return !exit})
}

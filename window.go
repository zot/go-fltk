package fltk

/*
#include "fltk.h"
*/
import "C"
import "unsafe"

type Window struct {
     Group
}

func NewWindow(w, h int) *Window {
	win := &Window{}
	win.Widget.Init(unsafe.Pointer(C.fltk_new_Window(C.int(w), C.int(h))))
	return win
}
func (w *Window) Show(args []string) {
     C.fltk_Window_show(w.ptr, C.int(0), unsafe.Pointer(uintptr(0)))
}

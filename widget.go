package fltk

/*
#include "fltk.h"
*/
import "C"
import "unsafe"

type Widget struct {
	Rectangle
	callback func()
	eventHandler func(*Event)
}

func emptyCallback() {}
func emptyHandler(*Event) {}

func (w *Widget) Init(p unsafe.Pointer) {
	w.Rectangle = Rectangle{(*C.Rectangle)(p)}
	w.callback = emptyCallback
	w.eventHandler = emptyHandler
	widgets[w.ptr] = w
}
func NewWidget(x, y, w, h int, text string) *Widget {
	wid := &Widget{}
	wid.Init(unsafe.Pointer(C.fltk_new_Widget(C.int(x), C.int(y), C.int(w), C.int(h), C.CString(text))))
	return wid
}
func (w *Widget) Box(box *C.Box) {C.fltk_Widget_box(w.ptr, box)}
func (w *Widget) LabelFont(font *C.Font) {C.fltk_Widget_labelfont(w.ptr, font)}
func (w *Widget) LabelSize(size int) {C.fltk_Widget_labelsize(w.ptr, C.int(size))}
func (w *Widget) LabelType(ltype *C.LabelType) {C.fltk_Widget_labeltype(w.ptr, ltype)}
func (w *Widget) SetCallback(f func()) {
	w.callback = f
	C.fltk_Widget_callback(w.ptr)
}
func (w *Widget) SetEventHandler(f func(*Event)) {w.eventHandler = f}
func (w *Widget) HandleCallback() {w.callback()}
func (w *Widget) HandleEvent(e *Event) {w.eventHandler(e)}

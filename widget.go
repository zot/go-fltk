package fltk

/*
#include "fltk.h"
*/
import "C"
import "unsafe"
import "fmt"

type Widget struct {
	ptr *C.Widget
	callback func()
	eventHandler func(*Event)
}

func emptyCallback() {}
func emptyHandler(*Event) {}

func NewWidget(x, y, w, h int, text... string) *Widget {
	return initWidget(&Widget{}, unsafe.Pointer(C.go_fltk_new_Widget(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
}
func (w *Widget) String() string {return fmt.Sprintf("Widget: %q", unsafe.Pointer(w.ptr))}
func (w *Widget) SetCallback(f func()) {
	w.callback = f
	C.go_fltk_Widget_callback(w.ptr)
}
func (w *Widget) ContinueEvent() bool {
	if C.go_fltk_event_stolen != 0 {
		C.go_fltk_Widget_continue_event(w.ptr)
		return true
	}
	return false
}
func (w *Widget) getWidget() *Widget {return w}
func (w *Widget) StealEvents(etypes int) {C.go_fltk_Widget_steal_events(w.ptr, C.int(etypes))}
func (w *Widget) Box(box *C.Box) {C.go_fltk_Widget_box(w.ptr, box)}
func (w *Widget) LabelFont(font *C.Font) {C.go_fltk_Widget_labelfont(w.ptr, font)}
func (w *Widget) LabelSize(size int) {C.go_fltk_Widget_labelsize(w.ptr, C.int(size))}
func (w *Widget) LabelType(ltype *C.LabelType) {C.go_fltk_Widget_labeltype(w.ptr, ltype)}
func (w *Widget) SetEventHandler(f func(*Event)) {w.eventHandler = f}
func (w *Widget) HandleCallback() {w.callback()}
func (w *Widget) HandleEvent(e *Event) {w.eventHandler(e)}
func (w *Widget) X() int {return int(C.go_fltk_Widget_x(w.ptr))}
func (w *Widget) Y() int {return int(C.go_fltk_Widget_y(w.ptr))}
func (w *Widget) W() int {return int(C.go_fltk_Widget_w(w.ptr))}
func (w *Widget) H() int {return int(C.go_fltk_Widget_h(w.ptr))}

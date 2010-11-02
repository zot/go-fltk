package fltk

/*
#include "fltk.h"
*/
import "C"
import "unsafe"

type Input struct {
     Widget
}

func NewInput(x, y, w, h int, text... string) *Input {
	var s *C.char

	if len(text) == 0 {
		s = (*C.char)(unsafe.Pointer(uintptr(0)))
	} else {
		s = C.CString(text[0])
	}
	i := &Input{}
	i.Widget.Init(unsafe.Pointer(C.fltk_new_Input(C.int(x), C.int(y), C.int(w), C.int(h), s)))
	return i
}
func (i *Input) StealEvents(etypes int) {C.fltk_Input_steal_events(i.ptr, C.int(etypes))}
func (i *Input) ContinueEvent() bool {
	if C.fltk_event_stolen != 0 {
		C.fltk_Input_continue_event(i.ptr)
		return true
	}
	return false
}

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
	i := &Input{}
	initWidget(i, unsafe.Pointer(C.go_fltk_new_Input(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return i
}

func (i *Input) GetText() string {return string(uintptr(unsafe.Pointer(C.go_fltk_Input_get_text(i.ptr))))}
func (i *Input) SetText(text string) bool {return C.go_fltk_Input_set_text(i.ptr, C.CString(text)) != C.int(0)}
func (i *Input) MousePosition() int {return int(C.go_fltk_Input_mouse_position(i.ptr))}
func (i *Input) GetPosition() int {return int(C.go_fltk_Input_get_position(i.ptr))}
func (i *Input) GetMark() int {return int(C.go_fltk_Input_get_mark(i.ptr))}

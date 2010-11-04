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

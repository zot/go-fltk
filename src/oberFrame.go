package fltk

/*
#include "fltk.h"
*/
import "C"
import "unsafe"

type OberFrame struct {
     Group
}

func NewOberFrame(x, y, w, h int, horizontal bool) *OberFrame {
	f := &OberFrame{}
	hor := 0
	if horizontal {hor = 1}
	initWidget(f, unsafe.Pointer(C.go_fltk_new_OberFrame(C.int(x), C.int(y), C.int(w), C.int(h), C.int(hor))))
	return f
}

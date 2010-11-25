package fltk

/*
#include "fltk.h"
*/
import "C"
import "unsafe"

type TextEditor struct {
     Widget
}

func NewTextEditor(x, y, w, h int, text... string) *TextEditor {
	i := &TextEditor{}
	initWidget(i, unsafe.Pointer(C.go_fltk_new_TextEditor(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return i
}

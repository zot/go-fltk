package fltk

/*
#include "fltk.h"
*/
import "C"
import "unsafe"

type Group struct {
     Widget
}

func NewPackedGroup(x, y, w, h int) *Group {
	win := &Group{}
	initWidget(win, unsafe.Pointer(C.go_fltk_new_PackedGroup(C.int(x), C.int(y), C.int(w), C.int(h))))
	return win
}

func (g *Group) Begin() {C.go_fltk_Group_begin(g.ptr)}
func (g *Group) End() {C.go_fltk_Group_end(g.ptr)}
func (g *Group) Add(w Widgety) {C.go_fltk_Group_add(g.ptr, w.getWidget().ptr)}
func (g *Group) Resizable(w Widgety) {C.go_fltk_Group_resizable(g.ptr, w.getWidget().ptr)}

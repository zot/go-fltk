package fltk

/*
#include "fltk.h"
*/
import "C"

type Group struct {
     Widget
}

func (g *Group) Begin() {C.go_fltk_Group_begin(g.ptr)}
func (g *Group) End() {C.go_fltk_Group_end(g.ptr)}
func (g *Group) Resizable(w Widgety) {C.go_fltk_Group_resizable(g.ptr, w.getWidget().ptr)}

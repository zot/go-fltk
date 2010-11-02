package fltk

/*
#include "fltk.h"
*/
import "C"

type Group struct {
     Widget
}

func (g *Group) Begin() {C.fltk_Group_begin(g.ptr)}
func (g *Group) End() {C.fltk_Group_end(g.ptr)}

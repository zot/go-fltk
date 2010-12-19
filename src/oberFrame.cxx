#include <fltk/Window.h>
#include <fltk/Group.h>
#include <fltk/PackedGroup.h>
#include <fltk/Widget.h>
#include <fltk/Input.h>
#include <fltk/TextEditor.h>
#include <fltk/run.h>
#include <fltk/events.h>
#include <fltk/Threads.h>
#include <fltk/draw.h>
#include "fltk.h"

static void assignLimits(Widget *ch, bool horizontal, int start, int end) {
	if (horizontal) {
		ch->x(start);
		ch->w(end - start);
	} else {
		ch->y(start);
		ch->h(end - start);
	}
}

void fltk::OberFrame::layout() {
	int end = horizontal ? w() : h();

	if (children() > 0) {
		assignLimits(child(0), true, child(0)->h(), w());
	}
	for (int i = children(); i-- > 2; ) {
		resizeFrame(dynamic_cast<fltk::OberFrame *>(child(i)), end);
	}
	if (children() > 1) {
		assignLimits(child(1), horizontal, 0, end);
	}
}

void fltk::OberFrame::resizeFrame(OberFrame *fr, int &end) {
	int start = horizontal ? fr->x() : fr->y();
	int min = fr->h();

	if (start > end - min) start = end - min;
	if (start < 0) start = 0;
	assignLimits(fr, horizontal, start, end);
	end = start;
}

void fltk::OberFrame::draw() {
	Group::draw();
	if (children() > 0) {
		Widget *ch = child(0);

		strokerect(0, 0, ch->h(), ch->h());
	}
}

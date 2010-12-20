/*
 * OberFrames have one widget at the top (like an input field) and either one widget or a set of
 * other ober frames underneath.  Ober frames draw a drag square at the upper left.  The
 * size of the drag square is equal to the height of the first child.  They lay out their first
 * child to the right of the drag square, stretched to the left side.  They lay out their other
 * children underneath the first child either left-to-right or top-to-bottom, depending on the
 * value of the "horizontal" field (given in the constructor).
 *
 * The user can drag the rectangle to resize the ober frame and it is constrained to be inside the
 * parent's children area (the parent should be an ober frame, as well).  Ober frames only display
 * their drag squares when they are inside other ober frames.
 *
 * If an ober frame has more than two children, the children other than the first must be
 * ober frames.
 */
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
#include <fltk/layout.h>
#include "fltk.h"
#include <stdio.h>
#include <typeinfo>

using namespace fltk;

static const char *tf(bool value) {
	return value ? "true" : "false";
}

OberFrame *go_fltk_new_OberFrame(int x, int y, int w, int h, int horizontal) {
	return new OberFrame(x, y, w, h, horizontal);
}

int OberFrame::handle(int evt) {
	int used = Group::handle(evt);
	int chh = child(0)->h();

	switch (evt) {
	case PUSH:
		if (!used) {
			pressed = true;
			redraw(dragRect());
			return true;
		}
		break;
	case RELEASE:
		pressed = false;
		redraw(dragRect());
		return true;
	case ENTER:
	case LEAVE:
	case MOVE: {
			bool inside = event_x() >= 0 && event_x() <= chh && event_y() >= 0 && event_y() <= chh;

			if (inside != in) {
				in = inside;
				if (!pressed) {
					redraw(dragRect());
				}
				return true;
			}
		}
		break;
	case DRAG:
		break;
	default:
		break;
	}
	return false;
}

static void assignLimits(Widget *ch, bool horizontal, int start, int end) {
	if (horizontal) {
		ch->x(start);
		ch->w(end - start);
	} else {
		ch->y(start);
		ch->h(end - start);
	}
}

void OberFrame::stretch(Widget *ch) {
	int tagSize = child(0)->h();

	if (horizontal) {
		ch->y(tagSize);
		ch->h(h() - tagSize);
	} else {
		ch->x(0);
		ch->w(w());
	}
	ch->layout_damage(LAYOUT_WH);
	ch->layout();
	ch->layout_damage(0);
}

void OberFrame::layout() {
	int end = horizontal ? w() : h();

	Widget::layout();
	if (children() > 2) {
		int minSize = 0;
		int dam = 0;

		for (int i = children(); i-- > 1; ) {
			OberFrame *ch = dynamic_cast<OberFrame*>(child(i));
			if (!ch) printf("CHILD %d IS NOT AN OBER FRAME, IT IS A %s!\n", i, typeid(child(i)).name());
			Widget *chch = ch->child(0);
			if (!chch) printf("CHILD %d HAS NO CHILD!\n", i);
			if (chch) minSize += chch->h();
		}
		if (horizontal) {
			if (w() < minSize)  {
				w(minSize);
				dam |= LAYOUT_H;
			}
		}
		minSize = (horizontal ? 0 : minSize) + child(0)->h();
		if (h() < minSize)  {
			h(minSize);
			dam |= LAYOUT_W;
		}
		if (dam) layout_damage(dam);
	}
	if (children() > 0) {
		assignLimits(child(0), true, showDragSquare() ? child(0)->h() : 0, w());
		child(0)->y(0);
		child(0)->layout();
		child(0)->layout_damage(0);
	}
	for (int i = children(); i-- > 2; ) {
		resizeFrame(dynamic_cast<OberFrame *>(child(i)), end);
	}
	if (children() > 1) {
		assignLimits(child(1), horizontal, horizontal ? 0 : child(0)->h(), end);
		stretch(child(1));
	}
	redraw(dragRect());
}

void OberFrame::resizeFrame(OberFrame *fr, int &end) {
	int minStart = horizontal ? 0 : child(0)->h();
	int start = horizontal ? fr->x() : fr->y();
	int minWid = fr->child(0)->h();

	if (start > end - minWid) start = end - minWid;
	if (start < minStart) start = minStart;
	assignLimits(fr, horizontal, start, end);
	end = start;
	stretch(fr);
}

void fltk::OberFrame::draw() {
	Group::draw();
	if (children() > 0 && showDragSquare()) {
		Widget *ch = child(0);
		int style = line_style();
		int width = line_width();
		const char *dashes = line_dashes();

		setcolor(pressed ? GREEN : in ? lerp(GREEN, WHITE, 0.75) : WHITE);
		fillrect(3, 3, ch->h() - 6, ch->h() - 6);
		setcolor(BLACK);
		line_style(SOLID, 2, 0);
		strokerect(2, 2, ch->h() - 4, ch->h() - 4);
		line_style(style, width, dashes);
	}
}

package fltk

/*
#include "fltk.h"
*/
import "C"
import "unsafe"

var UP_BOX *C.Box
var HELVETICA_BOLD_ITALIC *C.Font
var SHADOW_LABEL *C.LabelType

const (
	NO_EVENT	= iota
	PUSH
	RELEASE
	ENTER
	LEAVE
	DRAG
	FOCUS
	UNFOCUS
	KEY
	KEYUP
	FOCUS_CHANGE
	MOVE
	SHORTCUT
	DEACTIVATE
	ACTIVATE
	HIDE
	SHOW
	PASTE
	TIMEOUT
	MOUSEWHEEL
	DND_ENTER
	DND_DRAG
	DND_LEAVE
	DND_RELEASE
	TOOLTIP
)
const (
	NO_EVENT_MASK	= 1 << iota
	PUSH_MASK
	RELEASE_MASK
	ENTER_MASK
	LEAVE_MASK
	DRAG_MASK
	FOCUS_MASK
	UNFOCUS_MASK
	KEY_MASK
	KEYUP_MASK
	FOCUS_CHAN_MASKGE
	MOVE_MASK
	SHORTCUT_MASK
	DEACTIVATE_MASK
	ACTIVATE_MASK
	HIDE_MASK
	SHOW_MASK
	PASTE_MASK
	TIMEOUT_MASK
	MOUSEWHEEL_MASK
	DND_ENTER_MASK
	DND_DRAG_MASK
	DND_LEAVE_MASK
	DND_RELEASE_MASK
	TOOLTIP_MASK
)

type Event struct {
	Widget *Widget
	Event int
	State int
	Button int
	Clicks int
	Dx int
	Dy int
	Key int
	X int
	Y int
	XRoot int
	YRoot int
	Stolen bool
	TimeElapsed int
}

var widgets map[*C.Widget]*Widget

func init() {
	UP_BOX = C.fltk_get_UP_BOX()
	HELVETICA_BOLD_ITALIC = C.fltk_get_HELVETICA_BOLD_ITALIC()
	SHADOW_LABEL = C.fltk_get_SHADOW_LABEL()
	widgets = map[*C.Widget]*Widget{}
}
func Run() {
	C.fltk_run()
}
func Wait(time... float64) *Event {
	i := 0

	if len(time) == 0 {
		i = int(C.fltk_wait_forever())
	} else {
		i = int(C.fltk_wait(C.double(time[0])))
	}
	if C.fltk_callback_widget != (*C.Widget)(unsafe.Pointer(uintptr(0))) {
		widgets[C.fltk_callback_widget].HandleCallback()
	}
	if C.fltk_event_widget != (*C.Widget)(unsafe.Pointer(uintptr(0))) {
		evt := &Event{
			widgets[C.fltk_event_widget],
			int(C.fltk_event),
			int(C.fltk_event_state),
			int(C.fltk_event_button),
			int(C.fltk_event_clicks),
			int(C.fltk_event_dx),
			int(C.fltk_event_dy),
			int(C.fltk_event_key),
			int(C.fltk_event_x),
			int(C.fltk_event_y),
			int(C.fltk_event_x_root),
			int(C.fltk_event_y_root),
			int(C.fltk_event_stolen) != 0,
			i,
		}
		widgets[C.fltk_event_widget].HandleEvent(evt)
		return evt
	}
	return nil
}

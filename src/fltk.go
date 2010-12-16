package fltk

/*
#include "fltk.h"
*/
import "C"
import "unsafe"
//import "fmt"

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

/*! Flags returned by event_state(), and used as the high 16 bits
  of Widget::add_shortcut() values (the low 16 bits are all zero, so these
  may be or'd with key values).

  The inline function BUTTON(n) will turn n (1-8) into the flag for a
  mouse button.
*/
const (
  SHIFT		= 1 << 16 * iota	// 0x00010000,	/*!< Either shift key held down */
  CAPSLOCK						// 0x00020000,	/*!< Caps lock is toggled on */
  CTRL							// 0x00040000,	/*!< Either ctrl key held down */
  ALT							// 0x00080000,	/*!< Either alt key held down */
  NUMLOCK						// 0x00100000,	/*!< Num Lock turned on */
  _
  META							// 0x00400000,	/*!< "Windows" or the "Apple" keys held down */
  SCROLLLOCK					// 0x00800000,	/*!< Scroll Lock turned on */
  BUTTON1						// 0x01000000,	/*!< Left mouse button held down */
  BUTTON2						// 0x02000000,	/*!< Middle mouse button held down */
  BUTTON3						// 0x04000000,	/*!< Right mouse button held down */
  ANY_BUTTON = 0x7f000000		// 0x7f000000, /*!< Any mouse button (up to 8) */
  ACCELERATOR = C.ACCELERATOR	//!< ALT on Windows/Linux, CTRL on OS/X, use for menu accelerators
  OPTION = C.OPTION				//!< ALT|META on Windows/Linux, just ALT on OS/X, use as a drag modifier
  COMMAND = C.COMMAND			//!< CTRL on Windows/Linux, META on OS/X, use for menu shortcuts
)

type Event struct {
	Callback Widgety
	Widget Widgety
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
	Return int
//	TimeElapsed int
}

func (*Event) Continue() bool{
	if C.go_fltk_event_stolen != 0 {
		C.go_fltk_Event_continue()
		return true
	}
	return false
}

var widgets map[*C.Widget]Widgety

type Widgety interface {
	GetWidget() *Widget
	String() string
}

func debug(args... interface{}) {}
//func debug(args... interface{}) {fmt.Println(args...)}

func cStringOpt(s []string) *C.char {
	if len(s) == 0 {
		return (*C.char)(unsafe.Pointer(uintptr(0)))
	}
	return C.CString(s[0])
}

func init() {
	UP_BOX = C.go_fltk_get_UP_BOX()
	HELVETICA_BOLD_ITALIC = C.go_fltk_get_HELVETICA_BOLD_ITALIC()
	SHADOW_LABEL = C.go_fltk_get_SHADOW_LABEL()
	widgets = map[*C.Widget]Widgety{}
	Start()
}
func initWidget(w Widgety, p unsafe.Pointer) *Widget {
	w.GetWidget().ptr = (*C.Widget)(p)
	w.GetWidget().callback = emptyCallback
	w.GetWidget().eventHandler = emptyHandler
	widgets[w.GetWidget().ptr] = w
	return w.GetWidget()
}
func Start() {
	debug("started")
	go func() {
		C.go_fltk_run()
		debug("FLTK DONE")
	}()
}
func ReadEvent() *Event {
	debug("GO  WAITING FOR EVENT FROM FLTK")
	C.go_fltk_get_event()
	debug("GO GOT EVENT FROM FLTK")
	return &Event{
		widgets[C.go_fltk_callback_widget],
		widgets[C.go_fltk_event_widget],
		int(C.go_fltk_event),
		int(C.go_fltk_event_state),
		int(C.go_fltk_event_button),
		int(C.go_fltk_event_clicks),
		int(C.go_fltk_event_dx),
		int(C.go_fltk_event_dy),
		int(C.go_fltk_event_key),
		int(C.go_fltk_event_x),
		int(C.go_fltk_event_y),
		int(C.go_fltk_event_x_root),
		int(C.go_fltk_event_y_root),
		int(C.go_fltk_event_stolen) != 0,
		int(C.go_fltk_event_return),
	}
}
func ContinueEvent(used int) {
	debug("GO CONTINUING FLTK")
	C.go_fltk_continue_event(C.int(used))
}
func Handle(event *Event) int {
	if event.Callback != nil {
		debug("CALLBACK")
		event.Callback.GetWidget().HandleCallback()
	} else if event.Widget != nil {
		debug("-- widget:", event.Widget, "event:", event.Event)
		event.Widget.GetWidget().HandleEvent(event)
		debug("-- event returned: ", event.Return)
	}
	return event.Return
}
func Run(while func() bool) {
	for (while()) {ContinueEvent(Handle(ReadEvent()))}
}
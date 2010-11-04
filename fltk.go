package fltk

/*
#include "fltk.h"
*/
import "C"
import "unsafe"
import "fmt"

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

var widgets map[*C.Widget]Widgety

type Widgety interface {
	getWidget() *Widget
	String() string
}

func debug(args... interface{}) {}//{fmt.Println(args...)}

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
}
func initWidget(w Widgety, p unsafe.Pointer) *Widget {
	w.getWidget().ptr = (*C.Widget)(p)
	w.getWidget().callback = emptyCallback
	w.getWidget().eventHandler = emptyHandler
	widgets[w.getWidget().ptr] = w
	return w.getWidget()
}
/*
func Run() {
	C.go_fltk_run()
}
func Wait(time... float64) *Event {
	i := 0

	if len(time) == 0 {
		i = int(C.go_fltk_wait_forever())
	} else {
		i = int(C.go_fltk_wait(C.double(time[0])))
	}
	if C.go_fltk_callback_widget != (*C.Widget)(unsafe.Pointer(uintptr(0))) {
		widgets[C.go_fltk_callback_widget].getWidget().HandleCallback()
	}
	if C.go_fltk_event_widget != (*C.Widget)(unsafe.Pointer(uintptr(0))) {
		evt := &Event{
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
			i,
		}
		if evt.Widget != nil {
			debug("widget: " + evt.Widget.String())
			evt.Widget.getWidget().HandleEvent(evt)
		} else {
			debug("no widget for event")
		}
		return evt
	}
	return nil
}
*/
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
func ContinueEvent(i int) {
	debug("GO CONTINUING FLTK")
	C.go_fltk_continue_event(C.int(i))
}
func Handle(event *Event) int {
	if event.Callback != nil {
		debug("CALLBACK")
		event.Callback.getWidget().HandleCallback()
	} else if event.Widget != nil {
		event.Widget.getWidget().HandleEvent(event)
		fmt.Println("-- widget:", event.Widget, "event:", event.Event, " returned: ", event.Return)
	}
	return event.Return
}
func Run() {
	Start()
	for {
		ContinueEvent(Handle(ReadEvent()))
	}
}
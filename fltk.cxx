#include <fltk/Window.h>
#include <fltk/Widget.h>
#include <fltk/Input.h>
#include <fltk/run.h>
#include <fltk/events.h>
#include "fltk.h"
#include <stdio.h>

#define MAX_EVENT_TYPES 25

Widget *fltk_callback_widget = 0;
Widget *fltk_event_widget;
int fltk_event;
int fltk_event_state;
int fltk_event_button;
int fltk_event_clicks;
int fltk_event_dx;
int fltk_event_dy;
int fltk_event_key;
int fltk_event_x;
int fltk_event_y;
int fltk_event_x_root;
int fltk_event_y_root;
int fltk_event_stolen;

void set_event_fields() {
  fltk_event_widget = fltk::belowmouse();
  fltk_event = fltk::event();
  fltk_event_state = fltk::event_state();
  fltk_event_button = fltk::event_button();
  fltk_event_clicks = fltk::event_clicks();
  fltk_event_dx = fltk::event_dx();
  fltk_event_dy = fltk::event_dy();
  fltk_event_key = fltk::event_key();
  fltk_event_x = fltk::event_x();
  fltk_event_y = fltk::event_y();
  fltk_event_x_root = fltk::event_x_root();
  fltk_event_y_root = fltk::event_y_root();
}

class EventStealer {
public:
  int stealEvents;

  virtual int continue_event() = 0;

  EventStealer() : stealEvents(0) {}

  int handle(int event) {
    if ((stealEvents & (1 << event)) != 0) {
      fltk_event_stolen = 1;
      return 1;
    }
    return continue_event();
  }
};

class GInput : public EventStealer, public fltk::Input {
public:
  GInput(int x, int y, int w, int h, const char *label = 0) : EventStealer(), fltk::Input(x, y, w, h, label) {}

  int handle(int event) {return EventStealer::handle(event);}

  int continue_event() {return fltk::Input::handle(fltk::event());}
};

static void notify_callback(Widget *w, void *data) {
  fltk_callback_widget = w;
}

Box *fltk_get_UP_BOX() {return fltk::UP_BOX;}
Font *fltk_get_HELVETICA_BOLD_ITALIC() {return  fltk::HELVETICA_BOLD_ITALIC;}
LabelType *fltk_get_SHADOW_LABEL() {return fltk::SHADOW_LABEL;}
void fltk_run() {fltk::run();}
Window *fltk_new_Window(int w, int h) {return new Window(w, h);}
Widget *fltk_new_Widget(int x, int y, int w, int h, const char *text) {return new Widget(x, y, w, h, text);}
Input *fltk_new_Input(int x, int y, int w, int h, const char *text) {return new GInput(x, y, w, h, text);}
void fltk_Input_steal_events(Input *input, int events) {((GInput*)input)->stealEvents = events;}
void fltk_Input_continue_event(Input *input) {((GInput*)input)->continue_event();}
void fltk_Group_begin(Group *w) {w->begin();}
void fltk_Group_end(Group *w) {w->end();}
void fltk_Window_show(Window *w, int argc, void *argv) {w->show(argc, (char **)argv);}
void fltk_Widget_box(Widget *w, Box *box) {w->box(box);}
void fltk_Widget_callback(Widget *w) {w->callback(notify_callback);}
void fltk_Widget_labelfont(Widget *w, Font *font) {w->labelfont(font);}
void fltk_Widget_labelsize(Widget *w, int size) {w->labelsize(size);}
void fltk_Widget_labeltype(Widget *w, LabelType *type) {w->labeltype(type);}
int fltk_wait_forever() {
  fltk_callback_widget = 0;
  fltk_event_stolen = 0;
  int i = fltk::wait();

  set_event_fields();
  return i;
}
int fltk_wait(double time) {
  fltk_callback_widget = 0;
  fltk_event_stolen = 0;
  int i = fltk::wait(time);

  set_event_fields();
  return i;
}

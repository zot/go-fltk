#ifdef __cplusplus
#define TYPEDEF(type) typedef fltk::type type
namespace fltk {
class FL_API OberFrame : public Group {
public:
	bool horizontal;

	OberFrame(int x, int y, int w, int h, const char *label, bool horiz) : horizontal(horiz), Group(x, y, w, h, label) {}
	void draw();
	void layout();
	void resizeFrame(OberFrame *fr, int &end);
};
}
extern "C" {
#else
#define TYPEDEF(type) typedef struct type type
enum {
  CTRL		= 0x00040000,	/*!< Either ctrl key held down */
  ALT		= 0x00080000,	/*!< Either alt key held down */
  META		= 0x00400000,	/*!< "Windows" or the "Apple" keys held down */
#if defined(__APPLE__)
  ACCELERATOR	= CTRL,
  OPTION	= ALT,
  COMMAND	= META
#else
  ACCELERATOR	= ALT, //!< ALT on Windows/Linux, CTRL on OS/X, use for menu accelerators
  COMMAND	= CTRL,	//!< CTRL on Windows/Linux, META on OS/X, use for menu shortcuts
  OPTION	= ALT|META //!< ALT|META on Windows/Linux, just ALT on OS/X, use as a drag modifier
#endif
};
#endif

TYPEDEF(OberFrame);
TYPEDEF(Box);
TYPEDEF(Font);
TYPEDEF(LabelType);
TYPEDEF(Widget);
TYPEDEF(Group);
TYPEDEF(PackedGroup);
TYPEDEF(Window);
TYPEDEF(Input);
TYPEDEF(TextEditor);

extern Widget *go_fltk_callback_widget;
extern int go_fltk_event;
extern Widget *go_fltk_event_widget;
extern int go_fltk_event_state;
extern int go_fltk_event_button;
extern int go_fltk_event_clicks;
extern int go_fltk_event_dx;
extern int go_fltk_event_dy;
extern int go_fltk_event_key;
extern int go_fltk_event_x;
extern int go_fltk_event_y;
extern int go_fltk_event_x_root;
extern int go_fltk_event_y_root;
extern int go_fltk_event_stolen;
extern int go_fltk_event_return;

extern void free(void *ptr);
static inline void free_string(char* s) { free(s); }
static inline char* to_char_ptr(void *ptr) {return (char *)ptr;}

extern Box *go_fltk_get_UP_BOX();
extern Font *go_fltk_get_HELVETICA_BOLD_ITALIC();
extern LabelType *go_fltk_get_SHADOW_LABEL();
extern Window *go_fltk_new_Window(int w, int h);
extern PackedGroup *go_fltk_new_PackedGroup(int x, int y, int w, int h);
extern Widget *go_fltk_new_Widget(int x, int y, int w, int h, const char *text);
extern Input *go_fltk_new_Input(int x, int y, int w, int h, const char *text);
extern TextEditor *go_fltk_new_TextEditor(int x, int y, int w, int h, const char *text);
extern void go_fltk_Widget_steal_events(Widget *w, int events);
extern void go_fltk_Event_continue();
extern void go_fltk_Group_begin(Group *g);
extern void go_fltk_Group_end(Group *g);
extern void go_fltk_Group_add(Group *g, Widget *w);
extern void go_fltk_Group_resizable(Group *g, Widget *w);
extern void go_fltk_Widget_box(Widget *w, Box *box);
extern void go_fltk_Widget_labelfont(Widget *w, Font *font);
extern void go_fltk_Widget_labelsize(Widget *w, int size);
extern void go_fltk_Widget_labeltype(Widget *w, LabelType *type);
extern void go_fltk_Widget_callback(Widget *w);
extern void go_fltk_Window_destroy(Window *w);
extern void go_fltk_Window_show(Window *w, int argc, void *argv);
extern void go_fltk_Window_set_label(Window *w, const char *label);
extern int go_fltk_Widget_x(Widget *w);
extern int go_fltk_Widget_y(Widget *w);
extern int go_fltk_Widget_w(Widget *w);
extern int go_fltk_Widget_h(Widget *w);
extern void go_fltk_Widget_throw_focus(Widget *w);
extern const char *go_fltk_Input_get_text(Input *in);
extern int go_fltk_Input_set_text(Input *in, const char *text);
extern int go_fltk_Input_mouse_position(Input *in);
extern int go_fltk_Input_get_position(Input *in);
extern int go_fltk_Input_get_mark(Input *in);
extern void go_fltk_Input_set_position(Input *in, int p, int m);
extern void go_fltk_run();
extern void go_fltk_get_event();
extern void go_fltk_continue_event(int i);

#ifdef __cplusplus
}
#endif

#ifdef __cplusplus
#define FLTK(name) fltk::name
namespace fltk {
class FL_API OberFrame : public Group {
public:
	bool horizontal;
	bool pressed;
	bool in;

	OberFrame(int x, int y, int w, int h, bool horiz) : horizontal(horiz), pressed(false), in(false), Group(x, y, w, h, "") {}
	void draw();
	void layout();
	int handle(int);
	void resizeFrame(OberFrame *fr, int &end);
	void stretch(Widget *);
	Rectangle dragRect() const {return Rectangle(0, 0, child(0)->h(), child(0)->h());}
	bool showDragSquare() const {return dynamic_cast<OberFrame*>(parent());}
};
}
extern "C" {
#else
#define TYPEDEF(type) typedef struct type type
#define FLTK(name) name
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
#endif

extern FLTK(Widget) *go_fltk_callback_widget;
extern int go_fltk_event;
extern FLTK(Widget) *go_fltk_event_widget;
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

extern FLTK(Box) *go_fltk_get_UP_BOX();
extern FLTK(Font) *go_fltk_get_HELVETICA_BOLD_ITALIC();
extern FLTK(LabelType) *go_fltk_get_SHADOW_LABEL();
extern FLTK(Window) *go_fltk_new_Window(int w, int h);
extern FLTK(PackedGroup) *go_fltk_new_PackedGroup(int x, int y, int w, int h);
extern FLTK(Widget) *go_fltk_new_Widget(int x, int y, int w, int h, const char *text);
extern FLTK(Input) *go_fltk_new_Input(int x, int y, int w, int h, const char *text);
extern FLTK(TextEditor) *go_fltk_new_TextEditor(int x, int y, int w, int h, const char *text);
extern FLTK(OberFrame) *go_fltk_new_OberFrame(int x, int y, int w, int h, int horizontal);
extern void go_fltk_Widget_steal_events(FLTK(Widget) *w, int events);
extern void go_fltk_Event_continue();
extern void go_fltk_Group_begin(FLTK(Group) *g);
extern void go_fltk_Group_end(FLTK(Group) *g);
extern void go_fltk_Group_add(FLTK(Group) *g, FLTK(Widget) *w);
extern void go_fltk_Group_resizable(FLTK(Group) *g, FLTK(Widget) *w);
extern void go_fltk_Widget_box(FLTK(Widget) *w, FLTK(Box) *box);
extern void go_fltk_Widget_labelfont(FLTK(Widget) *w, FLTK(Font) *font);
extern void go_fltk_Widget_labelsize(FLTK(Widget) *w, int size);
extern void go_fltk_Widget_labeltype(FLTK(Widget) *w, FLTK(LabelType) *type);
extern void go_fltk_Widget_callback(FLTK(Widget) *w);
extern void go_fltk_Window_destroy(FLTK(Window) *w);
extern void go_fltk_Window_show(FLTK(Window) *w, int argc, void *argv);
extern void go_fltk_Window_set_label(FLTK(Window) *w, const char *label);
extern int go_fltk_Widget_x(FLTK(Widget) *w);
extern int go_fltk_Widget_y(FLTK(Widget) *w);
extern int go_fltk_Widget_w(FLTK(Widget) *w);
extern int go_fltk_Widget_h(FLTK(Widget) *w);
extern void go_fltk_Widget_throw_focus(FLTK(Widget) *w);
extern const char *go_fltk_Input_get_text(FLTK(Input) *in);
extern int go_fltk_Input_set_text(FLTK(Input) *in, const char *text);
extern int go_fltk_Input_mouse_position(FLTK(Input) *in);
extern int go_fltk_Input_get_position(FLTK(Input) *in);
extern int go_fltk_Input_get_mark(FLTK(Input) *in);
extern void go_fltk_Input_set_position(FLTK(Input) *in, int p, int m);
extern void go_fltk_run();
extern void go_fltk_get_event();
extern void go_fltk_continue_event(int i);

#ifdef __cplusplus
}
#endif

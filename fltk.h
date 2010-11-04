#ifdef __cplusplus
#define TYPEDEF(type) typedef fltk::type type
extern "C" {
#else
#define TYPEDEF(type) typedef struct type type
#endif

TYPEDEF(Box);
TYPEDEF(Font);
TYPEDEF(LabelType);
TYPEDEF(Widget);
TYPEDEF(Group);
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

extern void free(void *ptr);
static inline void free_string(char* s) { free(s); }

extern void go_fltk_init();
extern void go_fltk_run();
extern Box *go_fltk_get_UP_BOX();
extern Font *go_fltk_get_HELVETICA_BOLD_ITALIC();
extern LabelType *go_fltk_get_SHADOW_LABEL();
extern Window *go_fltk_new_Window(int w, int h);
extern Widget *go_fltk_new_Widget(int x, int y, int w, int h, const char *text);
extern Input *go_fltk_new_Input(int x, int y, int w, int h, const char *text);
extern TextEditor *go_fltk_new_TextEditor(int x, int y, int w, int h, const char *text);
extern void go_fltk_Widget_steal_events(Widget *w, int events);
extern void go_fltk_Widget_continue_event(Widget *w);
extern void go_fltk_Group_begin(Group *g);
extern void go_fltk_Group_end(Group *g);
extern void go_fltk_Group_resizable(Group *g, Widget *w);
extern void go_fltk_Widget_box(Widget *w, Box *box);
extern void go_fltk_Widget_labelfont(Widget *w, Font *font);
extern void go_fltk_Widget_labelsize(Widget *w, int size);
extern void go_fltk_Widget_labeltype(Widget *w, LabelType *type);
extern void go_fltk_Widget_callback(Widget *w);
extern void go_fltk_Window_show(Window *w, int argc, void *argv);
extern int go_fltk_Widget_x(Widget *w);
extern int go_fltk_Widget_y(Widget *w);
extern int go_fltk_Widget_w(Widget *w);
extern int go_fltk_Widget_h(Widget *w);
extern int go_fltk_wait_forever();
extern int go_fltk_wait(double time);
extern void go_fltk_get_event();
extern void go_fltk_continue_event();
extern void go_fltk_init();

#ifdef __cplusplus
}
#endif

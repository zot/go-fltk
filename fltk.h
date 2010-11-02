#ifdef __cplusplus
#define TYPEDEF(type) typedef fltk::type type
extern "C" {
#else
#define TYPEDEF(type) typedef struct type type
#endif

TYPEDEF(Box);
TYPEDEF(Font);
TYPEDEF(LabelType);
TYPEDEF(Rectangle);
TYPEDEF(Widget);
TYPEDEF(Group);
TYPEDEF(Window);
TYPEDEF(Input);

extern Widget *fltk_callback_widget;
extern int fltk_event;
extern Widget *fltk_event_widget;
extern int fltk_event_state;
extern int fltk_event_button;
extern int fltk_event_clicks;
extern int fltk_event_dx;
extern int fltk_event_dy;
extern int fltk_event_key;
extern int fltk_event_x;
extern int fltk_event_y;
extern int fltk_event_x_root;
extern int fltk_event_y_root;
extern int fltk_event_stolen;

extern void free(void *ptr);
static inline void free_string(char* s) { free(s); }

extern void fltk_init();
extern void fltk_run();
extern Box *fltk_get_UP_BOX();
extern Font *fltk_get_HELVETICA_BOLD_ITALIC();
extern LabelType *fltk_get_SHADOW_LABEL();
extern Window *fltk_new_Window(int w, int h);
extern Widget *fltk_new_Widget(int x, int y, int w, int h, const char *text);
extern Input *fltk_new_Input(int x, int y, int w, int h, const char *text);
extern void fltk_Input_steal_events(Input *input, int events);
extern void fltk_Input_continue_event(Input *input);
extern void fltk_Group_begin(Group *w);
extern void fltk_Group_end(Group *w);
extern void fltk_Widget_box(Widget *w, Box *box);
extern void fltk_Widget_labelfont(Widget *w, Font *font);
extern void fltk_Widget_labelsize(Widget *w, int size);
extern void fltk_Widget_labeltype(Widget *w, LabelType *type);
extern void fltk_Widget_callback(Widget *w);
extern void fltk_Window_show(Window *w, int argc, void *argv);
extern int fltk_wait_forever();
extern int fltk_wait(double time);

#ifdef __cplusplus
}
#endif

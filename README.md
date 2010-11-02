Go FLTK
=======

This is a simple go wrapper for [FLTK2](http://www.fltk.org/doc-2.0/html/index.html), which I did to support my Go version of Ober (based on Acme).  It's very small and should be fairly easy to extend to support whatever other widgets and functionality you need.  Please contribute changes you make.  The code is licensed under the very permissive ZLIB license.  FLTK2 is only somewhat supported, but I like its mouse-text behavior better than FLTK 1.3's.

This wrapper is far from complete, and I'm only planning on supporting what I need to do Ober, but I welcome additions to it.

EVENTS
======

Here's how I'm doing events.  When you call fltk.Wait() or fltk.Wait(), the toolkit will process at most one event and then call the event handler with an Event structure on the widget that received the event.  Normally, the event handler processes the event after the widget processes it, but you can "steal" events so that your event handler will get the event before the widget has a chance to process it.  If you want the widget to process a stolen event, just call widget.ContinueEvent().  Here are the two cases:

case 1, processing an event after the widget handles it (default case):

     user event --> widget --> go event handler

case 2, processing an event before the widget handles it (event stealing):

     user event --> go event handler [optionally: --> widget]

You install an event handler with widget.SetEventHandler() like this:

input.SetEventHandler(func(e fltk.Event){println(e.Event)}

if you want to steal events, call widget.StealEvents() with a bit mask for the FLTK event types.  There are constants for the event types and masks for them by appending _MASK to the name, like PUSH_MASK, which is just 1 << PUSH.  You call StealEvents() like this:

input.StealEvents((1 << PUSH_MASK) | (1 << DRAG_MASK))

There is a bool field in Event, called "Stolen" to tell you whether the event was stolen.  Here is an example of an event handler that prints a message for stolen push and drag events and then continues them:

	i.StealEvents(fltk.PUSH_MASK | fltk.DRAG_MASK)
	i.SetEventHandler(func(e fltk.Event) {
		if (e.Stolen) {
			println("CONTINUING EVENT: " + strconv.Itoa(e.Event))
			i.ContinueEvent()
		}
	})

There is an example program in the "examples" directory, called text.go that shows a text field and steals events like this.


DEPENDENCIES
============
The make file SHOULD auto-download and unpack FLTK2 for you.  When it builds the Go library, it statically links FLTK into the library so that the only dependencies are -lXext -lXinerama -lXft -lX11 -lXi -lm.

This only runs under X, because of the way I build it, but FLTK purports to run on windows and macs, so if you can test that and help with the Makefile, I'd appreciate it.


BUILDING
========
To build, run "make" in the top-level directory.  This assumes you have g++ installed, because FLTK is written in C++ and you need to use g++ to link the library so it can resolve mangled names.

To install the library, run "make install".

The example depends on the library being installed, so before building the example, you have to do make install.  To build the example, run "make examples".

If you change things in the library, sometimes it helps to run "make clean", because the dependencies for cgo_fltk.so don't always seem to work.

If you want to eliminate the fltk-library from your directory, you can run "make clean-fltk".


EXAMPLES
========
There are two examples: hello and text.  Hello is the Hello World program from the [FLTK docs](http://www.fltk.org/doc-2.0/html/index.html) and text is a window with a text field in it that steals the PUSH and DRAG events, printing a message and continuing the events.

FLTK=fltk-2.0.x-r7725
FLTK_DIST=/tmp/$(FLTK).tar.bz2
FLTK_URL=http://ftp.funet.fi/pub/mirrors/ftp.easysw.com/pub/fltk/snapshots/$(FLTK).tar.bz2
FLTK_LIB=$(FLTK)/lib/libfltk2.a

examples: install
	(cd examples; make)

# uncomment the next line to have 'make clean' remove the fltk-dist
#clean: clean-fltk

include $(GOROOT)/src/Make.inc

TARG=fltk
CGOFILES=\
	fltk.go \
	rectangle.go \
	widget.go \
	group.go \
	input.go \
	window.go
CGO_DEPS=fltk.o

CGO_LDFLAGS=fltk.o $(FLTK_LIB) -lXext -lXinerama -lXft -lX11 -lXi -lm

include $(GOROOT)/src/Make.pkg

fltk.o: fltk.h fltk.cxx $(FLTK_LIB)
	g++ -I$(FLTK) -c fltk.cxx

$(CGOTARG).so: $(GCC_OFILES) $(CGO_DEPS)
	g++ $(_CGO_CFLAGS_$(GOARCH)) -o $@ $(GCC_OFILES) $(CGO_LDFLAGS)  $(_CGO_LDFLAGS_$(GOOS))

$(FLTK_LIB): $(FLTK)/config.status
	(cd $(FLTK); make DIRS="src $(LOCALIMAGES) images")

$(FLTK)/config.status: $(FLTK)/stamp
	(cd $(FLTK); ./configure --disable-gl --disable-shared)

$(FLTK)/stamp: $(FLTK_DIST)
	tar xjf $(FLTK_DIST)
	touch $(FLTK)/stamp

$(FLTK_DIST):
	(cd /tmp; wget $(FLTK_URL))

clean-fltk:
	rm -rf $(FLTK)

clean: clean-local

clean-local:
	(cd examples; make clean)
	rm -f *~

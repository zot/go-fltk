BINS = text hello
pkgdir=$(GOROOT)/pkg/$(GOOS)_$(GOARCH)

all: $(BINS)

text hello : % : %.go $(pkgdir)/cgo_fltk.so $(pkgdir)/fltk.a
	8g $@.go && 8l -o $@ $@.8

clean:
	rm -f $(BINS) *.8

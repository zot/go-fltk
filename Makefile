all:
	cd src && $(MAKE) $(MFLAGS) all

install:
	cd src && $(MAKE) $(MFLAGS) install

examples:
	cd src && $(MAKE) $(MFLAGS) examples

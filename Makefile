# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

<<<<<<< HEAD:Makefile
all: test-sdl
=======
all: libs test-sdl
>>>>>>> 966c1ab204a35f6810dd89766dfd6f438248d6ea:Makefile

libs:
	make -C sdl install
	make -C ttf install

test-sdl: test-sdl.go libs
	$(GC) test-sdl.go
	$(LD) -o $@ test-sdl.$(O)

clean:
<<<<<<< HEAD:Makefile
	rm -f -r *.8 *.6 *.o */*.8 */*.6 */*.o */_obj test-sdl shoot.png
=======
	make -C sdl clean
	make -C ttf clean
	make -C 4s clean
	rm -f -r *.8 *.6 *.o */*.8 */*.6 */*.o */_obj test-sdl
>>>>>>> 966c1ab204a35f6810dd89766dfd6f438248d6ea:Makefile

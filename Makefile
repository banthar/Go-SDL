# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

all: install

install:
	make -C sdl install
	make -C ttf install
	make -C mixer install
	make -C gfx install

clean:
	make -C sdl clean
	make -C ttf clean
	make -C mixer clean
	make -C 4s clean
	make -C test clean
	make -C gfx clean

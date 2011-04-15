# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

.PHONY: all install clean

all:
	gomake -C sdl
	gomake -C ttf
	gomake -C mixer
	gomake -C gfx

install: all
	gomake -C sdl install
	gomake -C ttf install
	gomake -C mixer install
	gomake -C gfx install

clean:
	gomake -C sdl clean
	gomake -C ttf clean
	gomake -C mixer clean
	gomake -C test clean
	gomake -C gfx clean

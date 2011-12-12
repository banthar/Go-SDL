# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

all: install

install:
	gomake clean
	gomake -C sdl install
	gomake -C sdl/audio install
	gomake -C ttf install
	gomake -C mixer install
	gomake -C gfx install

clean:
	gomake -C sdl clean
	gomake -C sdl/audio clean
	gomake -C ttf clean
	gomake -C mixer clean
	gomake -C test clean
	gomake -C gfx clean
	gomake -C doc clean

nuke:
	gomake -C sdl nuke
	gomake -C sdl/audio nuke
	gomake -C ttf nuke
	gomake -C mixer nuke
	gomake -C gfx nuke

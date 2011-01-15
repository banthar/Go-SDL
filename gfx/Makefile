# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=sdl/gfx

GOFILES:=constants.go\
	framerate.go
CGOFILES:=rotozoom.go blitfunc.go imagefilter.go
CGO_LDFLAGS:=-lSDL_gfx -L/usr/local/lib
CGO_CFLAGS+=-I/usr/include/SDL
CLEANFILES+=gfx

include $(GOROOT)/src/Make.pkg

# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=github.com/banthar/Go-SDL/gfx

GOFILES:=constants.go\
	framerate.go

CLEANFILES+=gfx

GC=${O}g -I../sdl/_obj
LD=${O}l -L../sdl/_obj

include $(GOROOT)/src/Make.pkg

fmt:
	gofmt -w $(GOFILES)

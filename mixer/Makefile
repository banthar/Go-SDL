# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=sdl/mixer

GOFILES:=constants.go
CGOFILES:=mixer.go
CGO_CFLAGS:=-I/usr/local/include -Ic:/msys/1.0/local/include
CGO_LDFLAGS:=-lSDL_mixer -L/usr/local/lib -Lc:/msys/1.0/local/lib

CLEANFILES+=mixer

include $(GOROOT)/src/Make.pkg

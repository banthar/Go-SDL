# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=sdl/mixer

GOFILES:=constants.go
CGOFILES:=mixer.go chunks.go

ifeq ($(GOOS),freebsd)
CGO_CFLAGS:=-I/usr/local/include
CGO_LDFLAGS:=-lSDL_mixer -L/usr/local/lib
endif

CLEANFILES+=mixer

include $(GOROOT)/src/Make.pkg

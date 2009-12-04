# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=sdl/mixer

GOFILES:=mixer_defs.go
CGOFILES:=mixer.go
CGO_LDFLAGS:=-lSDL_mixer

CLEANFILES+=mixer

include $(GOROOT)/src/Make.pkg

# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=⚛sdl/mixer

GOFILES:=constants.go
CGOFILES:=mixer.go
CGO_LDFLAGS:=-lSDL_mixer

CLEANFILES+=mixer

include $(GOROOT)/src/Make.pkg

_cgo_defun.c _cgo_gotypes.go _cgo_export.c _cgo_export.h: $(CGOFILES)
	CGOPKGPATH=$(dir) cgo -- $(CGO_CFLAGS) $(CGOFILES)
	sed s/cgo_sdl/cgo_⚛sdl/ --in-place _cgo_defun.c

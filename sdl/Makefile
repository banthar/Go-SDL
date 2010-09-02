# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=⚛sdl

GOFILES:=\
	constants.$(O).go\
	event.go

# uncomment this line to enable exp/draw support
# GOFILES:=$(GOFILES) sdldraw.go

CGOFILES:=sdl.go
CGO_LDFLAGS:=`pkg-config --libs sdl` -lSDL_image
CGO_CFLAGS:=`pkg-config --cflags sdl`

include $(GOROOT)/src/Make.pkg

constants.$(O).go: constants.c
	godefs -g sdl constants.c > constants.$(O).go
	gofmt -w constants.$(O).go

_cgo_defun.c: $(CGOFILES)
	CGOPKGPATH=$(dir) cgo -- $(CGO_CFLAGS) $(CGOFILES)
	sed s/cgo_sdl/cgo_⚛sdl/ --in-place _cgo_defun.c

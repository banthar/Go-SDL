include $(GOROOT)/src/Make.inc

TARG=⚛sdl

GOFILES:=constants.go event.go structs.$(O).go

CGOFILES:=sdl.go

CGO_LDFLAGS:=`pkg-config --libs sdl` -lSDL_image

CGO_CFLAGS:=`pkg-config --cflags sdl`

CLEANFILES+=sdl

include $(GOROOT)/src/Make.pkg

constants.go: constants.c
	godefs -g sdl constants.c > constants.go
	gofmt -w constants.go

_cgo_defun.c: $(CGOFILES)
	CGOPKGPATH=$(dir) cgo -- $(CGO_CFLAGS) $(CGOFILES)
	sed s/cgo_sdl/cgo_⚛sdl/ --in-place _cgo_defun.c


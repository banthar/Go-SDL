include $(GOROOT)/src/Make.inc

TARG=⚛sdl/audio

CGOFILES=audio.go
CGO_LDFLAGS=callback.o `pkg-config --libs sdl`
CGO_CFLAGS=`pkg-config --cflags sdl`

include $(GOROOT)/src/Make.pkg

_cgo_defun.c: $(CGOFILES)
	CGOPKGPATH=$(dir) cgo -- $(CGO_CFLAGS) $(CGOFILES)
	sed s/cgo_sdl/cgo_⚛sdl/ --in-place _cgo_defun.c

_cgo_defun.c: callback.h


_obj/$(TARG).a: callback.o
_cgo_.so: callback.o

callback.o: callback.c
	$(HOST_CC) $(_CGO_CFLAGS_$(GOARCH)) -fPIC -O2 -o $@ -c $(CGO_CFLAGS) $*.c

callback.o: callback.h
